package main

import (
	"fmt"
	"weedy.com/publicChain/part1-Basic-Prototype/BLC"
)

func main() {

	block := BLC.NewBlock("Genenis Block", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	fmt.Println(block)

}
