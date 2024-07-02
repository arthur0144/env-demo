package main

import (
	"fmt"
	"os"
)

func main() {
	vita := os.Getenv("VITA")
	fmt.Println(vita)
}
