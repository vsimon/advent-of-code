package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const lookFor = 19690720

func runGravityAssist(ops []int, noun, verb int) int {
	opcodes := append(ops[:0:0], ops...)
	opcodes[1] = noun
	opcodes[2] = verb
	for i := 0; i < len(opcodes); i += 4 {
		if opcodes[i] == 1 {
			if opcodes[i+3] > len(opcodes) {
				return 0
			}
			opcodes[opcodes[i+3]] = opcodes[opcodes[i+2]] + opcodes[opcodes[i+1]]
		} else if opcodes[i] == 2 {
			if opcodes[i+3] > len(opcodes) {
				return 0
			}
			opcodes[opcodes[i+3]] = opcodes[opcodes[i+2]] * opcodes[opcodes[i+1]]
		} else if opcodes[i] == 99 {
			break
		} else {
			return 0
		}
	}
	return opcodes[0]
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	var ops []int
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	strOps := strings.Split(scanner.Text(), ",")
	for _, s := range strOps {
		op, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		ops = append(ops, op)
	}

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if lookFor == runGravityAssist(ops, noun, verb) {
				fmt.Println(100*noun + verb)
				return
			}
		}
	}

}
