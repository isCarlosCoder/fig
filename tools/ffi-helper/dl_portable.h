/*
 * dl_portable.h — Cross-platform dynamic loading abstraction for the FFI helper.
 *
 * POSIX (Linux, macOS): wraps dlopen / dlsym / dlclose / dlerror
 * Windows:              wraps LoadLibraryA / GetProcAddress / FreeLibrary / FormatMessage
 */
#ifndef DL_PORTABLE_H
#define DL_PORTABLE_H

#ifdef _WIN32

#include <windows.h>

typedef HMODULE dl_handle;

static inline dl_handle dl_open(const char* path) {
    return LoadLibraryA(path);
}

static inline void* dl_sym(dl_handle h, const char* name) {
    return (void*)GetProcAddress(h, name);
}

static inline void dl_close(dl_handle h) {
    FreeLibrary(h);
}

static inline const char* dl_error(void) {
    static char buf[512];
    DWORD err = GetLastError();
    if (err == 0) return "(no error)";
    FormatMessageA(
        FORMAT_MESSAGE_FROM_SYSTEM | FORMAT_MESSAGE_IGNORE_INSERTS,
        NULL, err, 0, buf, sizeof(buf), NULL);
    return buf;
}

#else /* POSIX (Linux, macOS, *BSD) */

#include <dlfcn.h>

typedef void* dl_handle;

static inline dl_handle dl_open(const char* path) {
    return dlopen(path, RTLD_LAZY);
}

static inline void* dl_sym(dl_handle h, const char* name) {
    return dlsym(h, name);
}

static inline void dl_close(dl_handle h) {
    dlclose(h);
}

static inline const char* dl_error(void) {
    const char* e = dlerror();
    return e ? e : "(no error)";
}

#endif /* _WIN32 */

#endif /* DL_PORTABLE_H */
