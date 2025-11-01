package intcode

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Execute(memory []int) []int {
	result := make([]int, len(memory))
	copy(result, memory)

	for i := 0; i < len(result); i += 4 {
		opcode := Opcode(result[i])
		switch opcode {
		case Add:
			p1 := result[i+1]
			p2 := result[i+2]
			p3 := result[i+3]

			log.Default().Printf("Adding %d and %d, storing at position %d\n", result[p1], result[p2], p3)
			result[p3] = result[p1] + result[p2]
		case Multiply:
			p1 := result[i+1]
			p2 := result[i+2]
			p3 := result[i+3]

			log.Default().Printf("Multiplying %d and %d, storing at position %d\n", result[p1], result[p2], p3)
			result[p3] = result[p1] * result[p2]
		case Halt:
			log.Default().Println("Halting execution")

			return result
		default:
			log.Fatalf("unknown opcode %d at position %d", opcode, i)
		}
	}

	return result
}

func ReadMemoryFromFile(filename string) []int {
	memory := []int{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()

		for _, s := range strings.Split(str, ",") {
			instruction, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			memory = append(memory, instruction)
		}
	}

	return memory
}

func MemoryEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
