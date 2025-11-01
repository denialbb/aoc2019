package intcode

import (
	"testing"
)

func runTestProgram(program_file string, expected_file string, t *testing.T) {
	program := ReadMemoryFromFile(program_file)
	expected := ReadMemoryFromFile(expected_file)
	results := Execute(program)

	if MemoryEquals(results, expected) == false {
		t.Errorf("Execute(%d) = %d; want %d", program, results, expected)
	}
}

func TestIntcode(t *testing.T) {
	runTestProgram("tests/program_1", "tests/expected_1", t)
	runTestProgram("tests/program_2", "tests/expected_2", t)
	runTestProgram("tests/program_3", "tests/expected_3", t)
	runTestProgram("tests/program_4", "tests/expected_4", t)
	runTestProgram("tests/program_5", "tests/expected_5", t)
}
