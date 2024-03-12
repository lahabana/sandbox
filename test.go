package main

import (
	"runtime"
)

func main() {
	println(runtime.NumCPU())

	println("Hello")
	println("::group:: my info")
	println("Some cool text")
	println("Evan more data")
	println("::endgroup::")
}
