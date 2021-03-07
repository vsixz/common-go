package hello

import (
	"fmt"
)

func Say(name string) {
	fmt.Printf("Hello, %s\n", name)
	fmt.Println("common-go version: v1.0.1")
}
