package main // package name

import (
	"fmt"

	"rsc.io/quote/v4"
)

func main() {
	fmt.Println("Hello, GO lang!")
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}

// search pkg https://pkg.go.dev
// import package
// download package: go mod tidy
