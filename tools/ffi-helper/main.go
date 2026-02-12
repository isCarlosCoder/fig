package main

/*
#cgo linux LDFLAGS: -ldl -rdynamic
#cgo darwin LDFLAGS: -ldl
#cgo windows LDFLAGS: -lkernel32
#include "dl_portable.h"
#include <stdlib.h>
#include <string.h>

// helper-exposed function that C libraries can call to invoke a registered Fig callback
// (call_cb_fn is implemented as a thin wrapper that forwards to the Go export call_cb_from_go)
char* call_cb_fn(const char* cbid, const char* arg);

// prototypes for wrapper functions provided in wrappers.c
// int wrappers
int ffi_call_int_fn0(void* fn);
int ffi_call_int_fn1(void* fn, int a);
int ffi_call_int_fn(void* fn, int a, int b);
int ffi_call_int_fn3(void* fn, int a, int b, int c);
// double wrappers
double ffi_call_double_fn0(void* fn);
double ffi_call_double_fn1(void* fn, double a);
double ffi_call_double_fn(void* fn, double a, double b);
// void wrappers
void ffi_call_void_fn0(void* fn);
void ffi_call_void_fn1_int(void* fn, int a);
void ffi_call_void_fn1_str(void* fn, char* a);
void ffi_call_void_fn2_str(void* fn, char* a, char* b);
// string wrappers
char* ffi_call_str_fn0(void* fn);
char* ffi_call_str_fn1(void* fn, char* a);
char* ffi_call_str_fn2(void* fn, char* a, char* b);
char* ffi_call_str_fn3(void* fn, char* a, int b, double c);
char* ffi_call_str_fn3_intint(void* fn, char* a, int b, int c);
char* ffi_call_str_fn3_strs(void* fn, char* a, char* b, char* c);
char* ffi_call_str_fn4_sisi(void* fn, char* a, int b, char* c, int d);
char* ffi_call_str_fn4_ssss(void* fn, char* a, char* b, char* c, char* d);
char* ffi_call_str_fn4_siid(void* fn, char* a, int b, int c, double d);

*/
import "C"

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// cbRequest used to send callback invocation requests from C into Fig via the helper
type cbRequest struct {
	cbid string
	args []interface{}
	resp chan map[string]interface{}
}

var cbReqCh = make(chan cbRequest, 16)
var cbPending = map[string]chan map[string]interface{}{}
var cbPendingMu sync.Mutex
var cbNext uint64

// memory managed by helper (for ownership semantics)
var memByID = map[string]unsafe.Pointer{}
var memSize = map[string]int{}
var memMu sync.Mutex
var memNext uint64

func serve(r io.Reader, w io.Writer) error {
	// notify on stderr that server started
	fmt.Fprintln(os.Stderr, "ffi-helper: server started")
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, "ffi-helper: panic:", r)
		}
	}()
	sc := bufio.NewScanner(r)
	enc := json.NewEncoder(w)
	var encMu sync.Mutex

	// request channel so we can read from the socket while a C call is blocking
	reqCh := make(chan map[string]interface{}, 16)

	// start a goroutine to read incoming messages and route responses to cbPending
	go func() {
		for sc.Scan() {
			var req map[string]interface{}
			if err := json.Unmarshal(sc.Bytes(), &req); err != nil {
				// log invalid JSON for debugging
				fmt.Fprintln(os.Stderr, "ffi-helper: invalid json:", string(sc.Bytes()), "err:", err)
				resp := map[string]interface{}{"ok": false, "error": "invalid json"}
				encMu.Lock()
				if err := enc.Encode(resp); err != nil {
					fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
					encMu.Unlock()
					return
				}
				encMu.Unlock()
				continue
			}
			// if message is a response to helper-initiated request (has 'ok' field), route to pending
			if _, isResp := req["ok"]; isResp {
				if idRaw, ok := req["id"]; ok {
					id := fmt.Sprintf("%v", idRaw)
					cbPendingMu.Lock()
					ch, has := cbPending[id]
					cbPendingMu.Unlock()
					if has {
						fmt.Fprintln(os.Stderr, "ffi-helper: got cb response id=", id, "resp=", req)
						select {
						case ch <- req:
						default:
						}
					}
				}
				continue
			}
			// otherwise, forward to the request processing goroutine
			reqCh <- req
		}
		close(reqCh)
	}()

	// callback request processing uses package-level channel and pending map
	go func() {
		for req := range cbReqCh {
			id := fmt.Sprintf("%d", atomic.AddUint64(&cbNext, 1))
			m := map[string]interface{}{"cmd": "invoke_callback", "cb": req.cbid, "args": req.args, "id": id}
			ch := make(chan map[string]interface{}, 1)
			cbPendingMu.Lock()
			cbPending[id] = ch
			cbPendingMu.Unlock()
			encMu.Lock()
			if err := enc.Encode(m); err != nil {
				fmt.Fprintln(os.Stderr, "ffi-helper: failed to send invoke_callback to fig:", err)
				ch <- map[string]interface{}{"ok": false, "error": err.Error()}
				encMu.Unlock()
			} else {
				fmt.Fprintln(os.Stderr, "ffi-helper: req invoke_callback id=", id)
				encMu.Unlock()
			}
			select {
			case r := <-ch:
				fmt.Fprintln(os.Stderr, "ffi-helper: cb goroutine received response id=", id, "resp=", r)
				req.resp <- r
			case <-time.After(3 * time.Second):
				req.resp <- map[string]interface{}{"ok": false, "error": "timeout"}
			}
			cbPendingMu.Lock()
			delete(cbPending, id)
			cbPendingMu.Unlock()
		}
	}()

	// in-memory state for load/sym/call
	var nextHandle int64 = 1

	symbols := map[int64]map[string]int64{}
	var nextSymbol int64 = 1
	ptrs := map[int64]unsafe.Pointer{}
	symbolByID := map[int64]unsafe.Pointer{}
	symbolTypeByID := map[int64]string{}

	for req := range reqCh {
		cmd, _ := req["cmd"].(string)
		id := req["id"]
		fmt.Fprintln(os.Stderr, "ffi-helper: req", cmd, "id=", id)
		switch cmd {
		case "ping":
			resp := map[string]interface{}{"ok": true, "resp": "pong"}
			if id != nil {
				resp["id"] = id
			}
			encMu.Lock()
			if err := enc.Encode(resp); err != nil {
				encMu.Unlock()
				fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
				return err
			}
			encMu.Unlock()
		case "sleep":
			msf, _ := req["ms"].(float64)
			ms := int(msf)
			fmt.Fprintln(os.Stderr, "ffi-helper: sleeping", ms, "ms")
			time.Sleep(time.Duration(ms) * time.Millisecond)
			resp := map[string]interface{}{"ok": true, "resp": "slept"}
			if id != nil {
				resp["id"] = id
			}
			encMu.Lock()
			if err := enc.Encode(resp); err != nil {
				encMu.Unlock()
				fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
				return err
			}
			encMu.Unlock()
		case "crash":
			fmt.Fprintln(os.Stderr, "ffi-helper: crashing as requested")
			os.Exit(1)
		case "load":
			// load via dlopen
			p, _ := req["path"].(string)
			if p == "" {
				resp := map[string]interface{}{"ok": false, "error": "missing path"}
				if id != nil {
					resp["id"] = id
				}
				encMu.Lock()
				enc.Encode(resp)
				encMu.Unlock()
				continue
			}
			cp := C.CString(p)
			defer C.free(unsafe.Pointer(cp))
			hdl := C.dl_open(cp)
			if hdl == nil {
				errstr := C.GoString(C.dl_error())
				resp := map[string]interface{}{"ok": false, "error": fmt.Sprintf("dlopen failed: %s", errstr)}
				if id != nil {
					resp["id"] = id
				}
				encMu.Lock()
				if err := enc.Encode(resp); err != nil {
					encMu.Unlock()
					fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
					return err
				}
				encMu.Unlock()
				continue
			}
			handle := nextHandle
			nextHandle++
			if symbols[handle] == nil {
				symbols[handle] = map[string]int64{}
			}
			ptrs[handle] = unsafe.Pointer(hdl)
			resp := map[string]interface{}{"ok": true, "handle": handle}
			if id != nil {
				resp["id"] = id
			}
			if err := enc.Encode(resp); err != nil {
				fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
				return err
			}
		case "sym":
			h, _ := req["handle"].(float64)
			name, _ := req["name"].(string)
			rtype, _ := req["rtype"].(string)
			if rtype == "" {
				rtype = "int"
			}
			hid := int64(h)
			if _, ok := symbols[hid]; !ok {
				resp := map[string]interface{}{"ok": false, "error": "invalid handle"}
				if id != nil {
					resp["id"] = id
				}
				encMu.Lock()
				if err := enc.Encode(resp); err != nil {
					encMu.Unlock()
					fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
					return err
				}
				encMu.Unlock()
				continue
			}
			cname := C.CString(name)
			defer C.free(unsafe.Pointer(cname))
			sym := C.dl_sym(C.dl_handle(ptrs[hid]), cname)
			if sym == nil {
				errstr := C.GoString(C.dl_error())
				resp := map[string]interface{}{"ok": false, "error": fmt.Sprintf("dlsym failed: %s", errstr)}
				if id != nil {
					resp["id"] = id
				}
				encMu.Lock()
				if err := enc.Encode(resp); err != nil {
					encMu.Unlock()
					fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
					return err
				}
				encMu.Unlock()
				continue
			}
			sid := nextSymbol
			nextSymbol++
			symbols[hid][name] = sid
			symbolByID[sid] = unsafe.Pointer(sym)
			symbolTypeByID[sid] = rtype
			resp := map[string]interface{}{"ok": true, "symbol": sid}
			if id != nil {
				resp["id"] = id
			}
			encMu.Lock()
			if err := enc.Encode(resp); err != nil {
				encMu.Unlock()
				fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
				return err
			}
			encMu.Unlock()
		case "call":
			// if a symbol id is provided, route to C function
			if symRaw, ok := req["symbol"]; ok {
				sid := int64(symRaw.(float64))
				ptr := symbolByID[sid]
				if ptr == nil {
					resp := map[string]interface{}{"ok": false, "error": "invalid symbol"}
					if id != nil {
						resp["id"] = id
					}
					encMu.Lock()
					if err := enc.Encode(resp); err != nil {
						encMu.Unlock()
						fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
						return err
					}
					encMu.Unlock()
					continue
				}
				rtype := symbolTypeByID[sid]
				args, _ := req["args"].([]interface{})
				switch rtype {
				case "int":
					if len(args) == 0 {
						res := C.ffi_call_int_fn0(ptr)
						resp := map[string]interface{}{"ok": true, "resp": float64(res)}
						if id != nil {
							resp["id"] = id
						}
						encMu.Lock()
						if err := enc.Encode(resp); err != nil {
							encMu.Unlock()
							fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
							return err
						}
						encMu.Unlock()
					} else if len(args) == 1 {
						a := C.int(int(args[0].(float64)))
						res := C.ffi_call_int_fn1(ptr, a)
						resp := map[string]interface{}{"ok": true, "resp": float64(res)}
						if id != nil {
							resp["id"] = id
						}
						encMu.Lock()
						if err := enc.Encode(resp); err != nil {
							encMu.Unlock()
							fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
							return err
						}
						encMu.Unlock()
					} else if len(args) == 2 {
						a := C.int(int(args[0].(float64)))
						b := C.int(int(args[1].(float64)))
						res := C.ffi_call_int_fn(ptr, a, b)
						resp := map[string]interface{}{"ok": true, "resp": float64(res)}
						if id != nil {
							resp["id"] = id
						}
						encMu.Lock()
						if err := enc.Encode(resp); err != nil {
							encMu.Unlock()
							fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
							return err
						}
						encMu.Unlock()
					} else if len(args) == 3 {
						a := C.int(int(args[0].(float64)))
						b := C.int(int(args[1].(float64)))
						c := C.int(int(args[2].(float64)))
						res := C.ffi_call_int_fn3(ptr, a, b, c)
						resp := map[string]interface{}{"ok": true, "resp": float64(res)}
						if id != nil {
							resp["id"] = id
						}
						encMu.Lock()
						if err := enc.Encode(resp); err != nil {
							encMu.Unlock()
							fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
							return err
						}
						encMu.Unlock()
					} else {
						resp := map[string]interface{}{"ok": false, "error": fmt.Sprintf("int call with %d args not supported", len(args))}
						if id != nil {
							resp["id"] = id
						}
						encMu.Lock()
						if err := enc.Encode(resp); err != nil {
							encMu.Unlock()
							fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
							return err
						}
						encMu.Unlock()
					}
				case "double":
					if len(args) == 0 {
						res := C.ffi_call_double_fn0(ptr)
						resp := map[string]interface{}{"ok": true, "resp": float64(res)}
						if id != nil {
							resp["id"] = id
						}
						encMu.Lock()
						if err := enc.Encode(resp); err != nil {
							encMu.Unlock()
							fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
							return err
						}
						encMu.Unlock()
					} else if len(args) == 1 {
						a := C.double(args[0].(float64))
						res := C.ffi_call_double_fn1(ptr, a)
						resp := map[string]interface{}{"ok": true, "resp": float64(res)}
						if id != nil {
							resp["id"] = id
						}
						encMu.Lock()
						if err := enc.Encode(resp); err != nil {
							encMu.Unlock()
							fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
							return err
						}
						encMu.Unlock()
					} else if len(args) == 2 {
						a := C.double(args[0].(float64))
						b := C.double(args[1].(float64))
						res := C.ffi_call_double_fn(ptr, a, b)
						resp := map[string]interface{}{"ok": true, "resp": float64(res)}
						if id != nil {
							resp["id"] = id
						}
						encMu.Lock()
						if err := enc.Encode(resp); err != nil {
							encMu.Unlock()
							fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
							return err
						}
						encMu.Unlock()
					} else {
						resp := map[string]interface{}{"ok": false, "error": fmt.Sprintf("double call with %d args not supported", len(args))}
						if id != nil {
							resp["id"] = id
						}
						encMu.Lock()
						if err := enc.Encode(resp); err != nil {
							encMu.Unlock()
							fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
							return err
						}
						encMu.Unlock()
					}
				case "void":
					// void return type — call function and return nil
					if len(args) == 0 {
						C.ffi_call_void_fn0(ptr)
					} else if len(args) == 1 {
						// check arg type
						if n, ok := args[0].(float64); ok {
							C.ffi_call_void_fn1_int(ptr, C.int(int(n)))
						} else if s, ok := args[0].(string); ok {
							cs := C.CString(s)
							C.ffi_call_void_fn1_str(ptr, cs)
							C.free(unsafe.Pointer(cs))
						} else {
							resp := map[string]interface{}{"ok": false, "error": "void: unsupported arg type"}
							if id != nil {
								resp["id"] = id
							}
							encMu.Lock()
							enc.Encode(resp)
							encMu.Unlock()
							continue
						}
					} else if len(args) == 2 {
						s0, ok0 := args[0].(string)
						s1, ok1 := args[1].(string)
						if ok0 && ok1 {
							cs0 := C.CString(s0)
							cs1 := C.CString(s1)
							C.ffi_call_void_fn2_str(ptr, cs0, cs1)
							C.free(unsafe.Pointer(cs0))
							C.free(unsafe.Pointer(cs1))
						} else {
							resp := map[string]interface{}{"ok": false, "error": "void: unsupported 2-arg combination"}
							if id != nil {
								resp["id"] = id
							}
							encMu.Lock()
							enc.Encode(resp)
							encMu.Unlock()
							continue
						}
					} else {
						resp := map[string]interface{}{"ok": false, "error": fmt.Sprintf("void call with %d args not supported", len(args))}
						if id != nil {
							resp["id"] = id
						}
						encMu.Lock()
						enc.Encode(resp)
						encMu.Unlock()
						continue
					}
					resp := map[string]interface{}{"ok": true, "resp": nil}
					if id != nil {
						resp["id"] = id
					}
					encMu.Lock()
					if err := enc.Encode(resp); err != nil {
						encMu.Unlock()
						fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
						return err
					}
					encMu.Unlock()
				case "string", "str":
					var toFree []*C.char
					defer func() {
						for _, p := range toFree {
							C.free(unsafe.Pointer(p))
						}
					}()
					// helper to extract a C string from a value (string, callback, mem ptr, or map→JSON)
					extract := func(x interface{}) (*C.char, error) {
						if s, ok := x.(string); ok {
							cs := C.CString(s)
							toFree = append(toFree, cs)
							return cs, nil
						}
						if m, ok := x.(map[string]interface{}); ok {
							if cbid, ok := m["__cb__"].(string); ok {
								cs := C.CString(cbid)
								toFree = append(toFree, cs)
								return cs, nil
							}
							if mid, ok := m["__mem__"].(string); ok {
								memMu.Lock()
								p, has := memByID[mid]
								memMu.Unlock()
								if has && p != nil {
									return (*C.char)(p), nil
								}
								return nil, fmt.Errorf("invalid memory id: %s", mid)
							}
							// generic map → JSON string
							js, err := json.Marshal(m)
							if err != nil {
								return nil, fmt.Errorf("cannot marshal object to json: %v", err)
							}
							cs := C.CString(string(js))
							toFree = append(toFree, cs)
							return cs, nil
						}
						return nil, fmt.Errorf("expected string or callback object")
					}

					// read arg_types metadata if provided
					var argTypesList []string
					if atRaw, ok := req["arg_types"].([]interface{}); ok {
						for _, v := range atRaw {
							if s, ok := v.(string); ok {
								argTypesList = append(argTypesList, s)
							}
						}
					}

					var cres *C.char
					// dispatch based on arg count and types
					if len(args) == 0 {
						cres = C.ffi_call_str_fn0(ptr)
					} else {
						// extract arg0 (always a string-like value for 1+ arg calls)
						arg0Str, err := extract(args[0])
						if err != nil {
							resp := map[string]interface{}{"ok": false, "error": err.Error()}
							if id != nil {
								resp["id"] = id
							}
							if err := enc.Encode(resp); err != nil {
								fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
								return err
							}
							continue
						}

						if len(args) == 1 {
							cres = C.ffi_call_str_fn1(ptr, arg0Str)
						} else if len(args) == 2 {
							a1Str, err := extract(args[1])
							if err != nil {
								resp := map[string]interface{}{"ok": false, "error": err.Error()}
								if id != nil {
									resp["id"] = id
								}
								if err := enc.Encode(resp); err != nil {
									fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
									return err
								}
								continue
							}
							cres = C.ffi_call_str_fn2(ptr, arg0Str, a1Str)
						} else if len(args) == 3 {
							if len(argTypesList) == 3 {
								t1 := argTypesList[1]
								t2 := argTypesList[2]
								isInt1 := t1 == "int" || t1 == "integer"
								isInt2 := t2 == "int" || t2 == "integer"
								isNum2 := t2 == "number" || t2 == "double" || t2 == "float"
								if isInt1 && isInt2 {
									n1 := int(args[1].(float64))
									n2 := int(args[2].(float64))
									cres = C.ffi_call_str_fn3_intint(ptr, arg0Str, C.int(n1), C.int(n2))
								} else if isInt1 && isNum2 {
									n1 := int(args[1].(float64))
									n2 := args[2].(float64)
									cres = C.ffi_call_str_fn3(ptr, arg0Str, C.int(n1), C.double(n2))
								} else {
									a1Str, err := extract(args[1])
									if err != nil {
										resp := map[string]interface{}{"ok": false, "error": err.Error()}
										if id != nil {
											resp["id"] = id
										}
										if err := enc.Encode(resp); err != nil {
											fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
											return err
										}
										continue
									}
									a2Str, err := extract(args[2])
									if err != nil {
										resp := map[string]interface{}{"ok": false, "error": err.Error()}
										if id != nil {
											resp["id"] = id
										}
										if err := enc.Encode(resp); err != nil {
											fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
											return err
										}
										continue
									}
									cres = C.ffi_call_str_fn3_strs(ptr, arg0Str, a1Str, a2Str)
								}
							} else {
								_, isStr1 := args[1].(string)
								_, isMap1 := args[1].(map[string]interface{})
								_, isStr2 := args[2].(string)
								_, isMap2 := args[2].(map[string]interface{})
								n1, isNum1 := args[1].(float64)
								n2, isNum2 := args[2].(float64)
								if (isStr1 || isMap1) && (isStr2 || isMap2) {
									a1Str, err := extract(args[1])
									if err != nil {
										resp := map[string]interface{}{"ok": false, "error": err.Error()}
										if id != nil {
											resp["id"] = id
										}
										if err := enc.Encode(resp); err != nil {
											fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
											return err
										}
										continue
									}
									a2Str, err := extract(args[2])
									if err != nil {
										resp := map[string]interface{}{"ok": false, "error": err.Error()}
										if id != nil {
											resp["id"] = id
										}
										if err := enc.Encode(resp); err != nil {
											fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
											return err
										}
										continue
									}
									cres = C.ffi_call_str_fn3_strs(ptr, arg0Str, a1Str, a2Str)
								} else if isNum1 && isNum2 {
									cres = C.ffi_call_str_fn3_intint(ptr, arg0Str, C.int(int(n1)), C.int(int(n2)))
								} else {
									resp := map[string]interface{}{"ok": false, "error": fmt.Sprintf("unsupported 3-arg string call: arg types %T, %T, %T", args[0], args[1], args[2])}
									if id != nil {
										resp["id"] = id
									}
									if err := enc.Encode(resp); err != nil {
										fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
										return err
									}
									continue
								}
							}
						} else if len(args) == 4 {
							// 4-arg string call: dispatch based on arg_types
							if len(argTypesList) == 4 {
								sig := ""
								for _, t := range argTypesList {
									switch {
									case t == "string" || t == "str":
										sig += "s"
									case t == "int" || t == "integer":
										sig += "i"
									case t == "double" || t == "float" || t == "number":
										sig += "d"
									default:
										sig += "?"
									}
								}
								switch sig {
								case "sisi":
									a2Str, err := extract(args[2])
									if err != nil {
										resp := map[string]interface{}{"ok": false, "error": err.Error()}
										if id != nil {
											resp["id"] = id
										}
										if err := enc.Encode(resp); err != nil {
											fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
											return err
										}
										continue
									}
									cres = C.ffi_call_str_fn4_sisi(ptr, arg0Str, C.int(int(args[1].(float64))), a2Str, C.int(int(args[3].(float64))))
								case "ssss":
									a1Str, err := extract(args[1])
									if err != nil {
										resp := map[string]interface{}{"ok": false, "error": err.Error()}
										if id != nil {
											resp["id"] = id
										}
										if err := enc.Encode(resp); err != nil {
											fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
											return err
										}
										continue
									}
									a2Str, err := extract(args[2])
									if err != nil {
										resp := map[string]interface{}{"ok": false, "error": err.Error()}
										if id != nil {
											resp["id"] = id
										}
										if err := enc.Encode(resp); err != nil {
											fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
											return err
										}
										continue
									}
									a3Str, err := extract(args[3])
									if err != nil {
										resp := map[string]interface{}{"ok": false, "error": err.Error()}
										if id != nil {
											resp["id"] = id
										}
										if err := enc.Encode(resp); err != nil {
											fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
											return err
										}
										continue
									}
									cres = C.ffi_call_str_fn4_ssss(ptr, arg0Str, a1Str, a2Str, a3Str)
								case "siid":
									cres = C.ffi_call_str_fn4_siid(ptr, arg0Str, C.int(int(args[1].(float64))), C.int(int(args[2].(float64))), C.double(args[3].(float64)))
								default:
									resp := map[string]interface{}{"ok": false, "error": fmt.Sprintf("unsupported 4-arg string signature: %s", sig)}
									if id != nil {
										resp["id"] = id
									}
									if err := enc.Encode(resp); err != nil {
										fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
										return err
									}
									continue
								}
							} else {
								resp := map[string]interface{}{"ok": false, "error": "4-arg string call requires arg_types metadata"}
								if id != nil {
									resp["id"] = id
								}
								if err := enc.Encode(resp); err != nil {
									fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
									return err
								}
								continue
							}
						} else {
							resp := map[string]interface{}{"ok": false, "error": fmt.Sprintf("string call with %d args not supported", len(args))}
							if id != nil {
								resp["id"] = id
							}
							if err := enc.Encode(resp); err != nil {
								fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
								return err
							}
							continue
						}
					} // end of 0-arg vs 1+ arg dispatch
					if cres == nil {
						resp := map[string]interface{}{"ok": false, "error": "call returned NULL"}
						if id != nil {
							resp["id"] = id
						}
						if err := enc.Encode(resp); err != nil {
							fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
							return err
						}
						continue
					}
					goStr := C.GoString(cres)
					C.free(unsafe.Pointer(cres))
					resp := map[string]interface{}{"ok": true, "resp": goStr}
					if id != nil {
						resp["id"] = id
					}
					if err := enc.Encode(resp); err != nil {
						fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
						return err
					}
				default:
					resp := map[string]interface{}{"ok": false, "error": "unsupported symbol type"}
					if id != nil {
						resp["id"] = id
					}
					if err := enc.Encode(resp); err != nil {
						fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
						return err
					}
				}
			} else {
				// fallback: sum/echo as before
				args, _ := req["args"].([]interface{})
				sum := float64(0)
				num := true
				for _, a := range args {
					switch v := a.(type) {
					case float64:
						sum += v
					default:
						num = false
					}
				}
				if num {
					resp := map[string]interface{}{"ok": true, "resp": sum}
					if id != nil {
						resp["id"] = id
					}
					if err := enc.Encode(resp); err != nil {
						fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
						return err
					}
				} else {
					resp := map[string]interface{}{"ok": true, "resp": args}
					if id != nil {
						resp["id"] = id
					}
					if err := enc.Encode(resp); err != nil {
						fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
						return err
					}
				}
			}
		case "alloc":
			szf, _ := req["size"].(float64)
			sz := int(szf)
			if sz <= 0 {
				resp := map[string]interface{}{"ok": false, "error": "invalid alloc size"}
				if id != nil {
					resp["id"] = id
				}
				encMu.Lock()
				if err := enc.Encode(resp); err != nil {
					encMu.Unlock()
					fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
					return err
				}
				encMu.Unlock()
				continue
			}
			ptr := C.malloc(C.size_t(sz))
			if ptr != nil {
				C.memset(ptr, 0, C.size_t(sz))
			}
			if ptr == nil {
				resp := map[string]interface{}{"ok": false, "error": "malloc failed"}
				if id != nil {
					resp["id"] = id
				}
				encMu.Lock()
				if err := enc.Encode(resp); err != nil {
					encMu.Unlock()
					fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
					return err
				}
				encMu.Unlock()
				continue
			}
			mid := fmt.Sprintf("m-%d", atomic.AddUint64(&memNext, 1))
			memMu.Lock()
			memByID[mid] = ptr
			memSize[mid] = sz
			memMu.Unlock()
			resp := map[string]interface{}{"ok": true, "mem_id": mid, "size": sz}
			if id != nil {
				resp["id"] = id
			}
			encMu.Lock()
			if err := enc.Encode(resp); err != nil {
				encMu.Unlock()
				fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
				return err
			}
			encMu.Unlock()
		case "strdup":
			// strdup(data) -> allocates C string copy, returns mem_id
			s, _ := req["data"].(string)
			if s == "" {
				resp := map[string]interface{}{"ok": false, "error": "strdup requires non-empty data string"}
				if id != nil {
					resp["id"] = id
				}
				encMu.Lock()
				if err := enc.Encode(resp); err != nil {
					encMu.Unlock()
					fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
					return err
				}
				encMu.Unlock()
				continue
			}
			cs := C.CString(s)
			sz := len(s) + 1 // includes NUL terminator
			mid := fmt.Sprintf("m-%d", atomic.AddUint64(&memNext, 1))
			memMu.Lock()
			memByID[mid] = unsafe.Pointer(cs)
			memSize[mid] = sz
			memMu.Unlock()
			resp := map[string]interface{}{"ok": true, "mem_id": mid, "size": sz}
			if id != nil {
				resp["id"] = id
			}
			encMu.Lock()
			if err := enc.Encode(resp); err != nil {
				encMu.Unlock()
				fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
				return err
			}
			encMu.Unlock()
		case "free":
			mid, _ := req["mem_id"].(string)
			if mid == "" {
				// fallback: try "id" for backwards compatibility
				mid, _ = req["id"].(string)
			}
			if mid == "" {
				resp := map[string]interface{}{"ok": false, "error": "free requires id"}
				if id != nil {
					resp["id"] = id
				}
				encMu.Lock()
				if err := enc.Encode(resp); err != nil {
					encMu.Unlock()
					fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
					return err
				}
				encMu.Unlock()
				continue
			}
			memMu.Lock()
			ptr, ok := memByID[mid]
			if ok {
				C.free(ptr)
				delete(memByID, mid)
				delete(memSize, mid)
			}
			memMu.Unlock()
			resp := map[string]interface{}{"ok": true}
			if id != nil {
				resp["id"] = id
			}
			encMu.Lock()
			if err := enc.Encode(resp); err != nil {
				encMu.Unlock()
				fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
				return err
			}
			encMu.Unlock()
		case "mem_write":
			mid, _ := req["mem_id"].(string)
			fmt.Fprintln(os.Stderr, "ffi-helper: mem_write requested mem_id=", mid)
			data, _ := req["data"].(map[string]interface{})
			b64, _ := data["__bytes__"].(string)
			offf, _ := req["offset"].(float64)
			off := int(offf)
			if mid == "" || b64 == "" {
				resp := map[string]interface{}{"ok": false, "error": "mem_write requires id and data"}
				if id != nil {
					resp["id"] = id
				}
				encMu.Lock()
				if err := enc.Encode(resp); err != nil {
					encMu.Unlock()
					fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
					return err
				}
				encMu.Unlock()
				continue
			}
			raw, err := base64.StdEncoding.DecodeString(b64)
			if err != nil {
				resp := map[string]interface{}{"ok": false, "error": "invalid base64"}
				if id != nil {
					resp["id"] = id
				}
				encMu.Lock()
				if err := enc.Encode(resp); err != nil {
					encMu.Unlock()
					fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
					return err
				}
				encMu.Unlock()
				continue
			}
			memMu.Lock()
			ptr, has := memByID[mid]
			sz, _ := memSize[mid]
			memMu.Unlock()
			if !has || ptr == nil {
				resp := map[string]interface{}{"ok": false, "error": "invalid mem id"}
				if id != nil {
					resp["id"] = id
				}
				encMu.Lock()
				if err := enc.Encode(resp); err != nil {
					encMu.Unlock()
					fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
					return err
				}
				encMu.Unlock()
				continue
			}
			if off < 0 || off+len(raw) > sz {
				resp := map[string]interface{}{"ok": false, "error": "write out of bounds"}
				if id != nil {
					resp["id"] = id
				}
				encMu.Lock()
				if err := enc.Encode(resp); err != nil {
					encMu.Unlock()
					fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
					return err
				}
				encMu.Unlock()
				continue
			}
			C.memcpy(unsafe.Pointer(uintptr(ptr)+uintptr(off)), unsafe.Pointer(&raw[0]), C.size_t(len(raw)))
			resp := map[string]interface{}{"ok": true}
			if id != nil {
				resp["id"] = id
			}
			encMu.Lock()
			if err := enc.Encode(resp); err != nil {
				encMu.Unlock()
				fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
				return err
			}
			encMu.Unlock()
		case "mem_read":
			mid, _ := req["mem_id"].(string)
			offf, _ := req["offset"].(float64)
			lnf, _ := req["len"].(float64)
			off := int(offf)
			ln := int(lnf)
			memMu.Lock()
			ptr, has := memByID[mid]
			sz, _ := memSize[mid]
			memMu.Unlock()
			if !has || ptr == nil {
				resp := map[string]interface{}{"ok": false, "error": "invalid mem id"}
				if id != nil {
					resp["id"] = id
				}
				encMu.Lock()
				if err := enc.Encode(resp); err != nil {
					encMu.Unlock()
					fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
					return err
				}
				encMu.Unlock()
				continue
			}
			if off < 0 || off+ln > sz {
				resp := map[string]interface{}{"ok": false, "error": "read out of bounds"}
				if id != nil {
					resp["id"] = id
				}
				encMu.Lock()
				if err := enc.Encode(resp); err != nil {
					encMu.Unlock()
					fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
					return err
				}
				encMu.Unlock()
				continue
			}
			buf := C.malloc(C.size_t(ln))
			C.memcpy(buf, unsafe.Pointer(uintptr(ptr)+uintptr(off)), C.size_t(ln))
			b := C.GoBytes(buf, C.int(ln))
			C.free(buf)
			encB64 := base64.StdEncoding.EncodeToString(b)
			resp := map[string]interface{}{"ok": true, "resp": map[string]interface{}{"__bytes__": encB64}}
			if id != nil {
				resp["id"] = id
			}
			encMu.Lock()
			if err := enc.Encode(resp); err != nil {
				encMu.Unlock()
				fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
				return err
			}
			encMu.Unlock()
		default:
			resp := map[string]interface{}{"ok": false, "error": "unknown command"}
			if id != nil {
				resp["id"] = id
			}
			if err := enc.Encode(resp); err != nil {
				fmt.Fprintln(os.Stderr, "ffi-helper: encode error:", err)
				return err
			}
		}
	}
	return sc.Err()
}

func serveSocket(sockPath string) error {
	// ensure parent exists
	_ = os.Remove(sockPath)
	ln, err := net.Listen("unix", sockPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = ln.Close()
		_ = os.Remove(sockPath)
	}()
	fmt.Fprintln(os.Stderr, "ffi-helper: socket server listening", sockPath)
	conn, err := ln.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()
	return serve(conn, conn)
}

//export call_cb_from_go
func call_cb_from_go(cbid *C.char, carg *C.char) *C.char {
	id := C.GoString(cbid)
	arg := ""
	if carg != nil {
		arg = C.GoString(carg)
	}
	rch := make(chan map[string]interface{}, 1)
	fmt.Fprintln(os.Stderr, "ffi-helper: call_cb_from_go received cb=", id, "arg=", arg)
	cbReqCh <- cbRequest{cbid: id, args: []interface{}{arg}, resp: rch}
	fmt.Fprintln(os.Stderr, "ffi-helper: call_cb_from_go queued request for cb=", id)
	resp := <-rch
	fmt.Fprintln(os.Stderr, "ffi-helper: call_cb_from_go got response for cb=", id, "resp=", resp)
	if ok, _ := resp["ok"].(bool); !ok {
		return nil
	}
	if s, ok := resp["resp"].(string); ok {
		return C.CString(s)
	}
	b, _ := json.Marshal(resp["resp"])
	return C.CString(string(b))
}

func main() {
	server := flag.Bool("server", false, "run as server reading JSON lines from stdin and responding")
	sock := flag.String("socket", "", "path to unix socket to listen on")
	flag.Parse()
	if *server {
		if err := serve(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, "server error:", err)
			os.Exit(1)
		}
		return
	}
	if *sock != "" {
		if err := serveSocket(*sock); err != nil {
			fmt.Fprintln(os.Stderr, "socket server error:", err)
			os.Exit(1)
		}
		return
	}
	fmt.Println("ffi-helper: stub helper (run with --server or --socket to act as RPC helper)")
}
