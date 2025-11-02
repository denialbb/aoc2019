package wires

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
)

type Point struct {
	x int
	y int
}

var centralPort Point = Point{x: 0, y: 0}

func FindClosestIntersection(intersections []Point) (Point, int) {
	distance := math.MaxInt
	closest_intersection := Point{}
	for _, p := range intersections {
		r := ManhattanDistance(centralPort, p)
		if r < distance {
			distance = r
			closest_intersection = p
		}
	}
	return closest_intersection, distance
}

func ManhattanDistance(pos1 Point, pos2 Point) int {
	return Abs((pos1.x - pos2.x)) + Abs((pos1.y - pos2.y))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ReadPathsFromFile(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(file)
	csvReader.Comma = ','

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

func GetWirePath(moves []string) map[Point]bool {
	path := make(map[Point]bool)
	x, y := 0, 0

	for _, move := range moves {
		direction := rune(move[0])
		distance, err := strconv.Atoi(move[1:])
		if err != nil {
			log.Fatal(err)
		}

		for range distance {
			switch direction {
			case 'R':
				x++
			case 'L':
				x--
			case 'U':
				y++
			case 'D':
				y--
			}
			path[Point{x: x, y: y}] = true
		}
	}

	return path
}

func Intersections(path1 map[Point]bool, path2 map[Point]bool) []Point {
	intersections := []Point{}

	for p := range path1 {
		if _, exists := path2[p]; exists {
			intersections = append(intersections, p)
		}
	}

	return intersections
}
