package main // package name

import (
	"example.com/greetings"
	"fmt"
	"log"
	//"rsc.io/quote/v4"
)

func main() {
	//fmt.Println("Hello, GO lang!")
	//fmt.Println(quote.Glass())
	//fmt.Println(quote.Go())
	//fmt.Println(quote.Hello())
	//fmt.Println(quote.Opt())

	//message := greetings.Hello("John")
	//fmt.Println(message)

	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("John")

	names := []string{"John", "Jane", "mike"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	fmt.Println(messages)

}

// # IMPORT A PACKAGE
// 1. search pkg https://pkg.go.dev
// 2. import package
// 3. download package: go mod tidy

// # IMPORT A LOCAL PACKAGE
// 1. go mod init {package-name}
// 2. cd {package-name}
// 3. setup page
// 4. cd ../{main package}
// 5. import "{namespace}"/{package-name}
// 6. edit {main-page}.mod: {namespace}/{package-name} => ../{package-name}
// 7. go mod tidy
