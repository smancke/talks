package main

/*
#include <stdio.h>
#include <stdlib.h>

extern void AGoFunc();
*/
import "C"
import "fmt"

//export AGoFunc
func AGoFunc() {
	fmt.Println("this is go code")
}
