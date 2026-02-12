#include <stdlib.h>
#include <string.h>

// helper exposes this symbol; the C test lib can call it to invoke a Fig callback
extern char* call_cb_fn(const char* cbid, const char* arg);

char* call_then_prefix(const char* cbid, const char* prefix) {
    char* res = call_cb_fn(cbid, "world");
    if (!res) return NULL;
    size_t lp = strlen(prefix);
    size_t lr = strlen(res);
    char* out = malloc(lp + lr + 1);
    if (!out) { free(res); return NULL; }
    memcpy(out, prefix, lp);
    memcpy(out + lp, res, lr);
    out[lp+lr] = '\0';
    free(res);
    return out;
}
