package main

import (
	"runtime"
)

func main() {
	println(runtime.NumCPU())

	println("Hello")
	println("::group::::error::my info")
	println("::group:: inside the group")
	println("Some cool text")
	println("::endgroup::")
	println("Some outside cool text")
	println("Evan more data")
	println("::endgroup::")
}
