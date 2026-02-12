/*
 * lib_asan_clean.c — FFI test library compiled with AddressSanitizer.
 * This library has NO memory errors — ASAN should produce zero warnings.
 * Used to verify that the ASAN integration pipeline works correctly.
 */
#include <stdlib.h>
#include <string.h>

/* Properly allocates, uses, and returns a string — no leaks, no overflows */
char* asan_safe_concat(const char* a, const char* b) {
    if (!a) a = "";
    if (!b) b = "";
    size_t la = strlen(a);
    size_t lb = strlen(b);
    char* out = (char*)malloc(la + lb + 1);
    if (!out) return NULL;
    memcpy(out, a, la);
    memcpy(out + la, b, lb);
    out[la + lb] = '\0';
    return out;
}

/* Simple int arithmetic — no memory involved */
int asan_safe_add(int a, int b) {
    return a + b;
}

/* Allocate and free properly within one call — no leak */
int asan_safe_strlen(const char* s) {
    if (!s) return 0;
    char* dup = strdup(s);
    int len = (int)strlen(dup);
    free(dup);
    return len;
}
