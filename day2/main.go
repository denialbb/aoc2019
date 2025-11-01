package intcode

import (
	"log"
)

func main() {
	memory := ReadMemoryFromFile("input")
	log.Default().Println("Initial memory:", memory)

	memory = Execute(memory)
	log.Default().Println("Final memory:", memory)
}
