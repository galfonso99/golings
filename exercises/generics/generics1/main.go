// generics1
// Make me compile!

package main

import (
    "fmt"
)
import "golang.org/x/exp/constraints"

func main() {
	print("Hello, World!")
	print(42)
}

func  print [T constraints.Integer|constraints.Float|~string] (value T) {
	fmt.Println(value)
}
