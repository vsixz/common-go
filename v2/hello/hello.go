package hello

import (
	"fmt"
)

func Say(firstName, lastName string) {
	fmt.Printf("Hello, %s %s\n", firstName,lastName)
	fmt.Println("common-go version: v2.0.0")
}
