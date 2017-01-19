package main

import "C"

func Operation(a, b int) int {
	return a + b
}

var OpName string = "add"
