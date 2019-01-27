package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var age int
var cmdLine = flag.NewFlagSet("question",flag.ExitOnError)

func init(){
	//flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	//flag.CommandLine.Usage = func() {
	//	fmt.Fprintf(os.Stderr,"Usage of %s:\n","executable file path")
	//	flag.PrintDefaults()
	//}
	cmdLine.StringVar(&name,"name","everyone","The greeting object.")
	cmdLine.IntVar(&age,"age",30,
		"The age of people")
}

func main(){
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("hello, %s! I'm %d years old.\n",name,age)
}