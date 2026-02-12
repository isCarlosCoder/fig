#include <stdlib.h>
#include <string.h>
#include <stdio.h>

// Simulates a C function that takes flattened nested struct fields.
// Conceptually:
//   struct Address { city: string, zip: int }
//   struct Person  { name: string, age: int, addr: Address }
// Flattened: person_with_addr(name: char*, age: int, city: char*, zip: int) -> char*
char* person_with_addr(char* name, int age, char* city, int zip) {
    if (!name) name = "";
    if (!city) city = "";
    char buf[512];
    int l = snprintf(buf, sizeof(buf), "%s age=%d city=%s zip=%d", name, age, city, zip);
    if (l < 0) return NULL;
    char* out = malloc(l + 1);
    if (!out) return NULL;
    memcpy(out, buf, l);
    out[l] = '\0';
    return out;
}
