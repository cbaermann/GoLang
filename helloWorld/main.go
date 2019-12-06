package main

//main is an executable package. must have a function called main associated with this.
//package == project == workspace
//first line must identify which it belongs to

import "fmt"

//go format, formats all the code in each file in the current directory
//gives package access to code in fmt(standard library package in go)
//some other packages: math, debut, encoding, crypto, io, debug
//golang.org/pkg for package list

func main() {
	fmt.Println("Hi there!")
}

//to run - go run filename.go

//go build - compiles program but does not execute

//go test, runs or executes any test files associated with current project
