package greetings

import "fmt"

func Hello(name string) string {
	//var message string
	//message = fmt.Sprintf("Hi, %v. Welcome to golang!", name)

	message := fmt.Sprintf("Hi, %v. Welcome to golang!", name)
	return message
}
