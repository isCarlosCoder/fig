#include <stdlib.h>
#include <string.h>
#include <stdio.h>

// === int wrappers ===
typedef int (*int_fn0)(void);
typedef int (*int_fn1)(int);
typedef int (*int_fn2)(int,int);
typedef int (*int_fn3)(int,int,int);
// mixed-type int wrappers
typedef int (*int_fn1_str)(const char*);
typedef int (*int_fn2_str_str)(const char*, const char*);
typedef int (*int_fn2_str_int)(const char*, int);
typedef int (*int_fn2_int_str)(int, const char*);
typedef int (*int_fn4_iisi)(int, int, const char*, int);

int ffi_call_int_fn0(void* fn) {
    return ((int_fn0)fn)();
}
int ffi_call_int_fn1(void* fn, int a) {
    return ((int_fn1)fn)(a);
}
int ffi_call_int_fn(void* fn, int a, int b) {
    return ((int_fn2)fn)(a,b);
}
int ffi_call_int_fn3(void* fn, int a, int b, int c) {
    return ((int_fn3)fn)(a,b,c);
}
int ffi_call_int_fn1_str(void* fn, char* a) {
    return ((int_fn1_str)fn)(a);
}
int ffi_call_int_fn2_str_str(void* fn, char* a, char* b) {
    return ((int_fn2_str_str)fn)(a, b);
}
int ffi_call_int_fn2_str_int(void* fn, char* a, int b) {
    return ((int_fn2_str_int)fn)(a, b);
}
int ffi_call_int_fn2_int_str(void* fn, int a, char* b) {
    return ((int_fn2_int_str)fn)(a, b);
}
int ffi_call_int_fn4_iisi(void* fn, int a, int b, char* c, int d) {
    return ((int_fn4_iisi)fn)(a, b, c, d);
}

// === double wrappers ===
typedef double (*double_fn0)(void);
typedef double (*double_fn1)(double);
typedef double (*double_fn2)(double,double);

double ffi_call_double_fn0(void* fn) {
    return ((double_fn0)fn)();
}
double ffi_call_double_fn1(void* fn, double a) {
    return ((double_fn1)fn)(a);
}
double ffi_call_double_fn(void* fn, double a, double b) {
    return ((double_fn2)fn)(a,b);
}

// === void wrappers ===
typedef void (*void_fn0)(void);
typedef void (*void_fn1_int)(int);
typedef void (*void_fn1_str)(char*);
typedef void (*void_fn2_str)(char*, char*);

void ffi_call_void_fn0(void* fn) {
    ((void_fn0)fn)();
}
void ffi_call_void_fn1_int(void* fn, int a) {
    ((void_fn1_int)fn)(a);
}
void ffi_call_void_fn1_str(void* fn, char* a) {
    ((void_fn1_str)fn)(a);
}
void ffi_call_void_fn2_str(void* fn, char* a, char* b) {
    ((void_fn2_str)fn)(a, b);
}

// === string wrappers ===
typedef char* (*str_fn0)(void);
typedef char* (*str_fn1)(char*);
typedef char* (*str_fn2)(char*,char*);
typedef char* (*str_fn2_intint)(int,int);
typedef char* (*str_fn3)(char*, int, double);
typedef char* (*str_fn3_intint)(char*, int, int);
typedef char* (*str_fn3s)(char*, char*, char*);
typedef char* (*str_fn4_sisi)(char*, int, char*, int);
typedef char* (*str_fn4_ssss)(char*, char*, char*, char*);
typedef char* (*str_fn4_siid)(char*, int, int, double);

char* ffi_call_str_fn0(void* fn) {
    return ((str_fn0)fn)();
}
char* ffi_call_str_fn1(void* fn, char* a) {
    return ((str_fn1)fn)(a);
}
char* ffi_call_str_fn2(void* fn, char* a, char* b) {
    return ((str_fn2)fn)(a,b);
}
char* ffi_call_str_fn2_intint(void* fn, int a, int b) {
    return ((str_fn2_intint)fn)(a,b);
}
char* ffi_call_str_fn3(void* fn, char* a, int b, double c) {
    return ((str_fn3)fn)(a,b,c);
}
char* ffi_call_str_fn3_intint(void* fn, char* a, int b, int c) {
    return ((str_fn3_intint)fn)(a,b,c);
}
char* ffi_call_str_fn3_strs(void* fn, char* a, char* b, char* c) {
    return ((str_fn3s)fn)(a,b,c);
}
char* ffi_call_str_fn4_sisi(void* fn, char* a, int b, char* c, int d) {
    return ((str_fn4_sisi)fn)(a,b,c,d);
}
char* ffi_call_str_fn4_ssss(void* fn, char* a, char* b, char* c, char* d) {
    return ((str_fn4_ssss)fn)(a,b,c,d);
}
char* ffi_call_str_fn4_siid(void* fn, char* a, int b, int c, double d) {
    return ((str_fn4_siid)fn)(a,b,c,d);
}

// forward to Go-exported function
extern char* call_cb_from_go(const char* cbid, const char* arg);

char* call_cb_fn(const char* cbid, const char* arg) {
    char* res = call_cb_from_go(cbid, arg);
    return res;
}
