package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	ops[1] = 12
	ops[2] = 2
	for i := 0; i < len(ops); i += 4 {
		if ops[i] == 1 {
			ops[ops[i+3]] = ops[ops[i+2]] + ops[ops[i+1]]
		} else if ops[i] == 2 {
			ops[ops[i+3]] = ops[ops[i+2]] * ops[ops[i+1]]
		} else if ops[i] == 99 {
			break
		} else {
			log.Fatal("Unknown op")
		}
	}
	fmt.Println(ops)
}
