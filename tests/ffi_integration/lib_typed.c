#include <stdlib.h>
#include <string.h>
#include <stdio.h>

// accepts char* name, int age, double score
char* person_summary(char* name, int age, double score) {
    if (!name) return NULL;
    fprintf(stderr, "person_summary called: age=%d score=%f\n", age, score);
    char buf[256];
    int l = snprintf(buf, sizeof(buf), "Name:%s Age:%d Score:%.2f", name, age, score);
    if (l < 0) return NULL;
    char* out = malloc(l + 1);
    if (!out) return NULL;
    memcpy(out, buf, l);
    out[l] = '\0';
    return out;
}

// variant that takes two ints instead of int/double
char* person_summary_ints(char* name, int age, int score) {
    if (!name) return NULL;
    fprintf(stderr, "person_summary_ints called: age=%d score=%d\n", age, score);
    char buf[256];
    int l = snprintf(buf, sizeof(buf), "Name:%s Age:%d Score:%d", name, age, score);
    if (l < 0) return NULL;
    char* out = malloc(l + 1);
    if (!out) return NULL;
    memcpy(out, buf, l);
    out[l] = '\0';
    return out;
}

char* person_summary_strs(char* name, char* ageStr, char* scoreStr) {
    if (!name) return NULL;
    int age = 0;
    double score = 0.0;
    if (ageStr) age = atoi(ageStr);
    if (scoreStr) score = atof(scoreStr);
    char buf[256];
    int l = snprintf(buf, sizeof(buf), "Name:%s Age:%d Score:%.2f", name, age, score);
    if (l < 0) return NULL;
    char* out = malloc(l + 1);
    if (!out) return NULL;
    memcpy(out, buf, l);
    out[l] = '\0';
    return out;
}
