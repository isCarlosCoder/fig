#include <stdlib.h>
#include <string.h>

// receives JSON string and prefix, returns prefix + JSON
char* struct_prefix(char* json, char* prefix) {
    if (!json || !prefix) return NULL;
    size_t lj = strlen(json);
    size_t lp = strlen(prefix);
    char* out = malloc(lp + lj + 1);
    if (!out) return NULL;
    memcpy(out, prefix, lp);
    memcpy(out + lp, json, lj);
    out[lp+lj] = '\0';
    return out;
}
