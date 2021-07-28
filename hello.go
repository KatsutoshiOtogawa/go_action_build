package main

import "fmt"

var REVISION string
var BUILD_TIME string

func main() {
	fmt.Println(fmt.Sprintf("Revision: %s,BUILD_TIME: %s", REVISION, BUILD_TIME))
	fmt.Printf("Hello World\n")
}
