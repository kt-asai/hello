package main

import (
	"fmt"

	"github.com/kt-asai/hello/sample"

	"rsc.io/quote/v3"
)

func main() {
	fmt.Println(quote.HelloV3())
	sample.ShowHelloSample()
}
