package tests

import "testing"

func useCrypto(code string) string {
return "use " + `"` + "crypto" + `"` + "; " + code
}

func TestCryptoHash(t *testing.T) {
out, err := runFig(t, useCrypto(`print(crypto.hash("hello"));`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "1.335831723e+09" && out != "1335831723" {
t.Fatalf("expected hash number, got %q", out)
}
}

func TestCryptoSha1(t *testing.T) {
out, err := runFig(t, useCrypto(`print(crypto.sha1("hello"));`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d" { t.Fatalf("wrong sha1, got %q", out) }
}

func TestCryptoSha256(t *testing.T) {
out, err := runFig(t, useCrypto(`print(crypto.sha256("hello"));`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824" { t.Fatalf("wrong sha256, got %q", out) }
}

func TestCryptoBase64Encode(t *testing.T) {
out, err := runFig(t, useCrypto(`print(crypto.base64Encode("FigLang"));`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "RmlnTGFuZw==" { t.Fatalf("expected 'RmlnTGFuZw==', got %q", out) }
}

func TestCryptoBase64Decode(t *testing.T) {
out, err := runFig(t, useCrypto(`print(crypto.base64Decode("RmlnTGFuZw=="));`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "FigLang" { t.Fatalf("expected 'FigLang', got %q", out) }
}

func TestCryptoBase64RoundTrip(t *testing.T) {
src := useCrypto(`let e = crypto.base64Encode("test123"); print(crypto.base64Decode(e));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "test123" { t.Fatalf("expected 'test123', got %q", out) }
}

func TestCryptoHexEncode(t *testing.T) {
out, err := runFig(t, useCrypto(`print(crypto.hexEncode("Fig"));`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "466967" { t.Fatalf("expected '466967', got %q", out) }
}

func TestCryptoHexDecode(t *testing.T) {
out, err := runFig(t, useCrypto(`print(crypto.hexDecode("466967"));`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "Fig" { t.Fatalf("expected 'Fig', got %q", out) }
}

func TestCryptoHexRoundTrip(t *testing.T) {
src := useCrypto(`let h = crypto.hexEncode("abc"); print(crypto.hexDecode(h));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "abc" { t.Fatalf("expected 'abc', got %q", out) }
}

func TestCryptoBase64BadInput(t *testing.T) {
_, err := runFig(t, useCrypto(`crypto.base64Decode("!!!invalid!!!");`))
if err == nil { t.Fatalf("expected error for invalid base64") }
}

func TestCryptoHexBadInput(t *testing.T) {
_, err := runFig(t, useCrypto(`crypto.hexDecode("ZZZZ");`))
if err == nil { t.Fatalf("expected error for invalid hex") }
}

func TestCryptoWrongType(t *testing.T) {
_, err := runFig(t, useCrypto(`crypto.sha256(123);`))
if err == nil { t.Fatalf("expected error for non-string argument") }
}
