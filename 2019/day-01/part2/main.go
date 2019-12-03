package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func fuelCalc(mass int64) int64 {
	var ret int64
	f := int64(math.Max(0, math.Floor(float64(mass/3.0))-2.0))
	for f > 0 {
		ret += f
		f = int64(math.Max(0, math.Floor(float64(f/3.0))-2.0))
	}
	return ret
}
func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	var fuel int64
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		mass, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		fuel += fuelCalc(mass)
	}

	fmt.Println(fuel)
}
