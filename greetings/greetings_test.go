package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "John"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("%v") = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

// go test -v
func TestWillPass(t *testing.T) {
	//var isFail bool = false // error
	isFail := true

	if isFail {
		t.Fatal("Failed")
	}
}
