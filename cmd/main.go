package main

import (
	"flag"
	echoer "gateway"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := echoer.RunHTTP(); err != nil {
		glog.Fatal(err)
	}
}
