package main

import (
	"flag"
	echoer "gateway"
)

func main() {
	flag.Parse()
	if err := echoer.RunHTTP(); err != nil {
		panic(err)
	}
}
