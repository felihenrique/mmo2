package main

import (
	"fmt"
	"os"
)

func main() {
	file := os.Getenv("GOFILE")

	if len(file) == 0 {
		fmt.Println("Env GOFILE is required")
		return
	}
}
