/*
 * asan_driver.c — Standalone ASAN test driver for FFI C libraries.
 * Compile with: gcc -fsanitize=address -g -O1 -o asan_driver asan_driver.c -L. -lasan_clean -ldl
 * Or dynamically: gcc -fsanitize=address -g -O1 -o asan_driver asan_driver.c -ldl
 *
 * This program loads a shared library, calls its functions, and any ASAN
 * instrumentation in the library (or driver) will detect memory errors.
 */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dlfcn.h>

int main(int argc, char** argv) {
    if (argc < 2) {
        fprintf(stderr, "usage: %s <path-to-lib.so>\n", argv[0]);
        return 1;
    }

    const char* libpath = argv[1];
    void* handle = dlopen(libpath, RTLD_LAZY);
    if (!handle) {
        fprintf(stderr, "dlopen failed: %s\n", dlerror());
        return 1;
    }

    /* Try asan_safe_add */
    typedef int (*add_fn)(int, int);
    add_fn safe_add = (add_fn)dlsym(handle, "asan_safe_add");
    if (safe_add) {
        int r = safe_add(10, 20);
        if (r != 30) {
            fprintf(stderr, "FAIL: asan_safe_add(10,20) = %d, expected 30\n", r);
            dlclose(handle);
            return 1;
        }
        printf("OK: asan_safe_add(10,20) = %d\n", r);
    }

    /* Try asan_safe_concat */
    typedef char* (*concat_fn)(const char*, const char*);
    concat_fn safe_concat = (concat_fn)dlsym(handle, "asan_safe_concat");
    if (safe_concat) {
        char* s = safe_concat("hello", " world");
        if (!s || strcmp(s, "hello world") != 0) {
            fprintf(stderr, "FAIL: asan_safe_concat mismatch\n");
            free(s);
            dlclose(handle);
            return 1;
        }
        printf("OK: asan_safe_concat = \"%s\"\n", s);
        free(s); /* properly free the returned string */
    }

    /* Try asan_safe_strlen */
    typedef int (*strlen_fn)(const char*);
    strlen_fn safe_strlen = (strlen_fn)dlsym(handle, "asan_safe_strlen");
    if (safe_strlen) {
        int len = safe_strlen("test");
        if (len != 4) {
            fprintf(stderr, "FAIL: asan_safe_strlen(\"test\") = %d, expected 4\n", len);
            dlclose(handle);
            return 1;
        }
        printf("OK: asan_safe_strlen(\"test\") = %d\n", len);
    }

    /* Try sum3 if available (from lib.c) */
    typedef int (*sum3_fn)(int, int, int);
    sum3_fn sum3 = (sum3_fn)dlsym(handle, "sum3");
    if (sum3) {
        int r = sum3(1, 2, 3);
        printf("OK: sum3(1,2,3) = %d\n", r);
    }

    /* Try dupstr if available */
    typedef char* (*dupstr_fn)(const char*);
    dupstr_fn dupstr = (dupstr_fn)dlsym(handle, "dupstr");
    if (dupstr) {
        char* s = dupstr("ASAN test");
        if (s) {
            printf("OK: dupstr = \"%s\"\n", s);
            free(s);
        }
    }

    dlclose(handle);
    printf("PASSED: All ASAN checks completed without errors.\n");
    return 0;
}
