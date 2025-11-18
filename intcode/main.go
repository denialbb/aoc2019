package main

import (
	"aoc2019/day2/intcode"
	"flag"
	"io"
	"log"
	"os"
)

var (
	verbose bool
	test bool
)

var debugLog *log.Logger

func execute(memory []int) {
	debugLog.Println("Initial memory:", memory)

	memory = intcode.Execute(memory)

	debugLog.Println("Final memory:", memory)
}

func findNounAndVerb(initialMemory []int, targetOutput int) (int, int) {
	return intcode.FindNounAndVerbForOutput(initialMemory, targetOutput)
}

func main() {
	var inputFilePath string

	flag.BoolVar(&verbose, "v", false, "enable verbose logging")
	flag.BoolVar(&test, "t", false, "TEST diagnostic program")
	flag.StringVar(&inputFilePath, "i", "input", "path to the input file")
	flag.Parse()

	if verbose {
		debugLog = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags)
	} else {
		debugLog = log.New(io.Discard, "", 0)
	}
	intcode.Init(debugLog)

	if test {
		intcode.TEST()
		return
	}

	if inputFilePath != "" {
		memory, err := intcode.ReadMemoryFromFile(inputFilePath)
		if err != nil {
			log.Fatal(err)
		}
		execute(*memory)
	}
}
