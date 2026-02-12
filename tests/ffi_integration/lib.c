#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int sum3(int a, int b, int c) {
    return a + b + c;
}

double mul2(double a, double b) {
    return a * b;
}

char* dupstr(const char* s) {
    if (!s) return NULL;
    return strdup(s);
}

char* concat(const char* a, const char* b) {
    if (!a) a = "";
    if (!b) b = "";
    size_t la = strlen(a);
    size_t lb = strlen(b);
    char *r = malloc(la + lb + 1);
    if (!r) return NULL;
    memcpy(r, a, la);
    memcpy(r + la, b, lb);
    r[la+lb] = '\0';
    return r;
}

char* join3(const char* a, const char* b, const char* c) {
    if (!a) a = "";
    if (!b) b = "";
    if (!c) c = "";
    size_t la = strlen(a);
    size_t lb = strlen(b);
    size_t lc = strlen(c);
    char *r = malloc(la + lb + lc + 1);
    if (!r) return NULL;
    memcpy(r, a, la);
    memcpy(r + la, b, lb);
    memcpy(r + la + lb, c, lc);
    r[la+lb+lc] = '\0';
    return r;
}
