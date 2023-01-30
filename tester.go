package tester

import (
	"fmt"
	module "github.com/YerkhanG/import_test"
)

func main() {
	// Get a greeting message and print it.
	message := module.Hello("Gladys")
	fmt.Println(message)
}
