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
)

func main() {
	flag.BoolVar(&verbose, "v", false, "enable verbose logging")
	flag.Parse()

	var debugLog *log.Logger
	if verbose {
		debugLog = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags)
	} else {
		debugLog = log.New(io.Discard, "", 0)
	}

	intcode.Init(debugLog)

	debugLog.Println("This will only show if -v is set")

	memory := intcode.ReadMemoryFromFile("input")
	memory = intcode.Restore1202ProgramAlarm(memory)

	debugLog.Println("Initial memory:", memory)

	memory = intcode.Execute(memory)

	debugLog.Println("Final memory:", memory)

	log.Default().Printf("Value at position 0: %d\n", memory[0])
}
