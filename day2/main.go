package main

import (
	"aoc2019/day2/intcode"
	"log"
)

func main() {
	memory := intcode.ReadMemoryFromFile("input")
	log.Default().Println("Initial memory:", memory)

	memory = intcode.Execute(memory)
	log.Default().Println("Final memory:", memory)
}
