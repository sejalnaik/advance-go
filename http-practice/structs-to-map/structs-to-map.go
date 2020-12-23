package main

import (
	"fmt"

	"github.com/fatih/structs"
)

type Person struct {
	Name   string
	Age    int
	IsMale bool
}

func main() {
	sejal := &Person{
		Name:   "sejal",
		Age:    26,
		IsMale: false,
	}
	sejalMap := structs.Map(sejal)

	fmt.Println(sejalMap["name"])
	fmt.Println(sejalMap["age"])
	fmt.Println(sejalMap["isMale"])
}
