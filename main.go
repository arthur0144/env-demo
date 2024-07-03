package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	op1 := os.Getenv("OP1")
	op2 := os.Getenv("OP2")
	operation := os.Getenv("OPERATION")

	op1v, err := strconv.Atoi(op1)
	if err != nil {
		fmt.Println("Не преобразовали op1", err)
		return
	}

	op2v, err := strconv.Atoi(op2)
	if err != nil {
		fmt.Println("Не преобразовали op2", err)
		return
	}

	if operation == "add" {
		fmt.Println(op1v + op2v)
	} else if operation == "mul" {
		fmt.Println(op1v * op2v)
	} else {
		fmt.Println("Неизвестная операция.")
	}
}
