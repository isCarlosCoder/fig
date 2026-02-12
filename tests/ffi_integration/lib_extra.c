#include <stdlib.h>
#include <string.h>
#include <stdio.h>

// global counter for void side-effect testing
static int g_counter = 0;

// void, 0-arg: increments global counter
void inc_counter(void) {
    g_counter++;
}

// void, 1-arg int: sets global counter to given value
void set_counter(int v) {
    g_counter = v;
}

// int, 0-arg: returns current counter value
int get_counter(void) {
    return g_counter;
}

// int, 1-arg: returns n * 2
int double_int(int n) {
    return n * 2;
}

// int, 2-arg: returns a + b 
int add2(int a, int b) {
    return a + b;
}

// double, 0-arg: returns pi approximation
double get_pi(void) {
    return 3.14159265;
}

// double, 1-arg: returns n * n
double square(double n) {
    return n * n;
}

// string, 0-arg: returns static greeting
char* hello(void) {
    return strdup("hello from C");
}

// string, 1-arg: returns uppercased first char + rest
char* shout(const char* s) {
    if (!s) return NULL;
    size_t len = strlen(s);
    char* out = malloc(len + 2);
    if (!out) return NULL;
    memcpy(out, s, len);
    out[len] = '!';
    out[len + 1] = '\0';
    return out;
}

// void, 1-arg string: prints to stderr (side-effect observable)
void print_msg(char* msg) {
    if (msg) fprintf(stderr, "print_msg: %s\n", msg);
}
