/*
The C language version with same generated output is smaller 83004 Bytes
compared with golang 2068291 Bytes, both on linux kernal Ubuntun specifically,
the compile and execution time are about the same. 

#include <stdio.h>

int main() {
	printf("Hello main program. This is not golang.\n");
	return 0;
}
*/

 
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello main program. This is not golang.")
}

