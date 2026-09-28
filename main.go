package main

import (
	"fmt"

	"IR_JVM/generate"
)

func main() {
	fmt.Println(generate.CONST{Type: 0, Const: generate.INT{10}})
}
