#include <stdlib.h>
#include <string.h>

// expects pointer to a null-terminated string
char* echo_mem(char* p) {
    if (!p) return NULL;
    size_t l = strlen(p);
    char* out = malloc(l + 1);
    if (!out) return NULL;
    memcpy(out, p, l);
    out[l] = '\0';
    return out;
}

// prefix the content of memory with prefix string
char* prefix_mem(char* memptr, char* prefix) {
    if (!memptr || !prefix) return NULL;
    size_t lm = strlen(memptr);
    size_t lp = strlen(prefix);
    char* out = malloc(lm + lp + 1);
    if (!out) return NULL;
    memcpy(out, prefix, lp);
    memcpy(out + lp, memptr, lm);
    out[lp+lm] = '\0';
    return out;
}
