package intcode

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestIntcode(t *testing.T) {
	testFile, err := os.Open("tests/test")
	if err != nil {
		t.Fatalf("failed to open test file: %s", err)
	}
	defer testFile.Close()

	resultsFile, err := os.Open("tests/results")
	if err != nil {
		t.Fatalf("failed to open results file: %s", err)
	}
	defer resultsFile.Close()

	testScanner := bufio.NewScanner(testFile)
	resultsScanner := bufio.NewScanner(resultsFile)

	program := []int{}
	expected := []int{}
	results := []int{}

	for testScanner.Scan() {
		if !resultsScanner.Scan() {
			t.Fatal("results file has fewer lines than test file")
		}

		instruction, err := strconv.Atoi(testScanner.Text())
		if err != nil {
			t.Fatalf("failed to parse program: %s", err)
		}

		expected_instruction, err := strconv.Atoi(resultsScanner.Text())
		if err != nil {
			t.Fatalf("failed to parse results: %s", err)
		}

		program = append(program, instruction)
		expected = append(expected, expected_instruction)
	}

	results = Execute(program)

	if MemoryEquals(results, expected) == false {
		t.Errorf("Execute(%d) = %d; want %d", program, results, expected)
	}

	if resultsScanner.Scan() {
		t.Fatal("test file has fewer lines than results file")
	}

	if err := testScanner.Err(); err != nil {
		t.Fatalf("error scanning test file: %s", err)
	}

	if err := resultsScanner.Err(); err != nil {
		t.Fatalf("error scanning results file: %s", err)
	}

}
