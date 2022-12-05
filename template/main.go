package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(solve())
}

func solve() (answer string) {
	input := readInputFile()

	return input[0]
}

func readInputFile() (input []string) {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return
}
