package main

import (
	"demo/lib"
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "Everyone", "get someone's name")
}

func main() {
	flag.Parse()
	lib4.SayHello(name)
}
