package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[0:] {
		s += sep + "[" + strconv.Itoa(i) + ":" + arg + "]"
		sep = " "
	}
	fmt.Println(s)
}
