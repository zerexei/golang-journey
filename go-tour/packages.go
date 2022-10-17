// By convention, the package name is the same as the last element of the import path.
// For instance, the "math/rand" package comprises files that begin with the statement package rand.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}


// import types

// v1
import "fmt"
import "math"
// v2

import (
	"fmt"
	"math"
)