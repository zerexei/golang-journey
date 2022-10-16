package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	//var message string
	//message = fmt.Sprintf("Hi, %v. Welcome to golang!", name)
	//message := fmt.Sprintf("Hi, %v. Welcome to golang!", name)

	if name == "" {
		return "", errors.New("empty name")
	}

	//var message string = fmt.Sprintf("Hi, %v. Welcome to golang!", name)
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// ? why go doesn't tell that the variable isn't defined
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

// Go executes `init` functions automatically at program startup,
// after global variables have been initialized.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome to golang!",
		"Great to see you, %v!",
		"Hail, %v!, Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
