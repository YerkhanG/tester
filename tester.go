package tester

import (
	"fmt"
	"github.com/YerkhanG/Golang-/tree/master/module"
)

func main() {
	// Get a greeting message and print it.
	message := module.Hello("Gladys")
	fmt.Println(message)
}
