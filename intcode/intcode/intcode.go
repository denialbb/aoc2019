package intcode

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"math"
)

var debugLog *log.Logger //nolint:gochecknoglobals

func Init(logger *log.Logger) {
	debugLog = logger
}

func Execute(memory []int) []int {
	_memory := make([]int, len(memory))
	copy(_memory, memory)
	// debugLog.Println("Memory:", _memory)

	// I guess I mean #inputs + #parameters
	var nparams int
	for i := 0; i < len(_memory); i += nparams+1 {
		code := _memory[i]
		opcode,modes := parseCode(code)
		switch opcode {
		case ADD:
			nparams = 3
			debugState(_memory, i, nparams)
			if len(modes) < nparams {
				for i := len(modes); i < nparams -1; i++ {
					modes = append(modes, 0) 
				}
				modes = append(modes, 1) // we write here
			}
			debugLog.Printf("ADD, modes: %v", modes)
			params := getParameters(_memory, i, modes, nparams)
			debugLog.Printf("memory[%d] = %d + %d\n",
				params[2],
				params[0],
				params[1])

			_memory[params[2]] = params[0] + params[1]

		case MULT:
			nparams = 3
			debugState(_memory, i, nparams)
			if len(modes) < nparams {
				for i := len(modes); i < nparams -1; i++ {
					modes = append(modes, 0) 
				}
				modes = append(modes, 1) // we write here
			}
			debugLog.Printf("MULT, modes: %v", modes)
			params := getParameters(_memory, i, modes, nparams)
			debugLog.Printf("memory[%d] = %d * %d\n",
				params[2],
				params[0],
				params[1])

			_memory[params[2]] = params[0] * params[1]

		case INPUT:
			nparams = 1
			debugState(_memory, i, nparams)
			debugLog.Printf("INPUT, modes: %v", modes)
			var input int
			print("> ")
			fmt.Scan(&input)
			out_ptr := getValue(_memory, i, 1)
			debugLog.Printf("memory[%d] = %d\n", out_ptr, input)
			_memory[out_ptr] = input

		case OUT:
			nparams = 1
			debugState(_memory, i, nparams)
			debugLog.Printf("OUT, modes: %v", modes)
			params := getParameters(_memory, i, modes, nparams)
			debugLog.Printf("OUTPUT: %d\n", params[0])
			println(params[0])
			
		case HALT:
			debugLog.Println("HALT received, stopping execution")
			return _memory

		default:
			log.Fatalf("unknown opcode %d at position %d", opcode, i)
		}
	}

	return _memory
}

func getStar(memory []int, index int, parameter int) int {
	return getParameter(memory, index, parameter, Position)
}

func getValue(memory []int, index int, parameter int) int {
	return getParameter(memory, index, parameter, Immediate)
}

func getParameter(memory []int, index int, parameter int, mode Mode) int {
	var res int

	switch mode {
	case Position:
		res = memory[memory[index+parameter]]
	case Immediate:
		res = memory[index+parameter]
	default:
		log.Fatalf("unknown mode %d", mode)
	}

	return res
}
func getParameters(memory []int, index int, modes []Mode, n_parameters int) []int {
	if len(modes) > n_parameters {
		debugLog.Printf("number of modes: %d, number of parameters: %d", 
			len(modes), n_parameters)
		log.Fatal("number of modes must less or equal to parameters")
	}
	params := make([]int, n_parameters)
	for i := range n_parameters {
		if i < len(modes) {
			switch modes[i] {
			case Position:
				params[i] = memory[memory[index+i+1]]
			case Immediate:
				params[i] = memory[index+i+1]
			default:
				log.Fatalf("unknown mode %d", modes[i])
			}
		} else { // omitted modes are Position
			params[i] = memory[memory[index+i+1]]
		}
	}
	// debugLog.Println(params)
	return params
}

func parseCode(code int) (Opcode,[]Mode) {
	length := intLength(code)
	modes := make([]Mode,0)

	// first 2 digits are opcode
	opcode := Opcode(code%100)
	// rest of the digits are modes for each parameter
	// leftmost 0s are omitted
	var leftmost int
	for i := range length-2 {
		leftmost = code/int(math.Pow10(i+2))
		modes = append(modes, Mode((leftmost)%10))
	}
	
	return opcode, modes
}

func intLength(n int) int {
	if n == 0 {
		return 1
	}

	count := 0
	for n != 0 {
		n /= 10
		count++
	}

	return count
}

func ReadMemoryFromFile(filename string) (*[]int, error) {
	memory := []int{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()

		for _, s := range strings.Split(str, ",") {
			instruction, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			memory = append(memory, instruction)
		}
	}

	return &memory, nil
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

func debugState(memory []int, index int, scope int) {
	memory_slice := make([]int, 0)
	for i := range scope+1 {
		memory_slice = append(memory_slice, memory[index+i])
	}
	debugLog.Print("[",index,"]: ",memory_slice)
}
