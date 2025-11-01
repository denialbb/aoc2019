package intcode

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Execute(memory []int) []int {
	result := make([]int, len(memory))

	return result
}

func ReadMemoryFromFile(filename string) []int {
	memory := []int{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		instruction, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		memory = append(memory, instruction)
	}

	return memory
}

func MemoryEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		 a[i] != b[i] {
			return false
		}
	}
	return true
}
