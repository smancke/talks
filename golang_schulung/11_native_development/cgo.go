package main

/*
#include <stdio.h>
#include <stdlib.h>

extern void AGoFunc();

void myprint(char* s) {
	printf("%s\n", s);
	AGoFunc();
}
*/
import "C"

import "unsafe"

func main() {
	cs := C.CString("Hello from stdio")
	C.myprint(cs)
	C.printf(cs)
	C.free(unsafe.Pointer(cs))
}
