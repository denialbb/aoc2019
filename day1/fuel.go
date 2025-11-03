package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func recursiveFuel(mass int) int {
	fuel := getFuel(mass)
	if fuel <= 0 {
		return 0
	}
	return fuel + recursiveFuel(fuel)
}

func getFuel(mass int) int {
	if mass <= 0 {
		return 0
	}

	fuel_mass := (mass / 3) - 2

	if fuel_mass > 0 {
		return fuel_mass
	} else {
		return 0
	}
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
		module_mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		result += recursiveFuel(module_mass)
	}

	fmt.Println(result)
}
