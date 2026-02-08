package builtins

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash/fnv"

	"github.com/iscarloscoder/fig/environment"
)

func init() {
	register(newModule("crypto",
		// hash(s) — returns a 32-bit FNV-1a hash as integer
		fn("hash", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("hash() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("hash() argument must be a string")
			}
			h := fnv.New32a()
			h.Write([]byte(args[0].Str))
			return environment.NewNumber(float64(h.Sum32())), nil
		}),

		// sha1(s) — returns SHA-1 hex string
		fn("sha1", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("sha1() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("sha1() argument must be a string")
			}
			sum := sha1.Sum([]byte(args[0].Str))
			return environment.NewString(hex.EncodeToString(sum[:])), nil
		}),

		// sha256(s) — returns SHA-256 hex string
		fn("sha256", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("sha256() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("sha256() argument must be a string")
			}
			sum := sha256.Sum256([]byte(args[0].Str))
			return environment.NewString(hex.EncodeToString(sum[:])), nil
		}),

		// base64Encode(s) — encodes string to base64
		fn("base64Encode", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("base64Encode() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("base64Encode() argument must be a string")
			}
			encoded := base64.StdEncoding.EncodeToString([]byte(args[0].Str))
			return environment.NewString(encoded), nil
		}),

		// base64Decode(s) — decodes base64 string
		fn("base64Decode", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("base64Decode() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("base64Decode() argument must be a string")
			}
			decoded, err := base64.StdEncoding.DecodeString(args[0].Str)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("base64Decode() invalid base64: %v", err)
			}
			return environment.NewString(string(decoded)), nil
		}),

		// hexEncode(s) — encodes string to hex
		fn("hexEncode", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("hexEncode() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("hexEncode() argument must be a string")
			}
			return environment.NewString(hex.EncodeToString([]byte(args[0].Str))), nil
		}),

		// hexDecode(s) — decodes hex string
		fn("hexDecode", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("hexDecode() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("hexDecode() argument must be a string")
			}
			decoded, err := hex.DecodeString(args[0].Str)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("hexDecode() invalid hex: %v", err)
			}
			return environment.NewString(string(decoded)), nil
		}),
	))
}
