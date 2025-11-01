package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFuel(mass int) int {
	return (mass / 3) - 2

}

func main() {
	input_file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer input_file.Close()

	filescanner := bufio.NewScanner(input_file)

	result := 0
	for filescanner.Scan() {
		line := filescanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		result += getFuel(i)
	}

	fmt.Println(result)
}
