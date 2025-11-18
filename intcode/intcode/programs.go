package intcode

import (
	"fmt"
	"log"
)

func Restore1202ProgramAlarm(memory []int) []int {
	debugLog.Println("Restoring 1202 program alarm state")
	memory[1] = 12
	memory[2] = 2

	return memory
}

const (
	minNoun = 0
	maxNoun = 99
	minVerb = 0
	maxVerb = 99
)

func FindNounAndVerbForOutput(initialMemory []int, targetOutput int) (int, int) {
	debugLog.Printf("Finding noun and verb for target output %d\n", targetOutput)

	for noun := minNoun; noun <= maxNoun; noun++ {
		for verb := minVerb; verb <= maxVerb; verb++ {
			memory := make([]int, len(initialMemory))
			copy(memory, initialMemory)
			memory[1] = noun
			memory[2] = verb
			result := Execute(memory)[0]

			if result == targetOutput {
				debugLog.Printf("Found noun %d and verb %d for target output %d\n", noun, verb, targetOutput)
				return noun, verb
			}
		}
	}

	return -1, -1
}

// TEST diagnostic program
func TEST() { 
	var system_ID ID
	print("Diagnostic Test...\nEnter system ID: ")
	fmt.Scan(&system_ID)

	switch system_ID {
	case ACU:
		println("Running diagnostics for ACU (re-enter the id when prompted)...")
		program_file := "./programs/diagnostic.i"
		program, err := ReadMemoryFromFile(program_file)
		if err != nil {
			log.Fatalf("Error reading program file: %v", err)
		}
		Execute(*program)
	case TRC:
		println("Running diagnostics for TRC (re-enter the id when prompted)...")
		program_file := "./programs/diagnostic.i"
		program, err := ReadMemoryFromFile(program_file)
		if err != nil {
			log.Fatalf("Error reading program file: %v", err)
		}
		Execute(*program)

	default:
		log.Fatalf("Unknown id %d", system_ID)

	}

	println("DONE")
}
