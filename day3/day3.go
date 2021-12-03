package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day3.demoinput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // for every line in the file

	}
	if err != nil {
		log.Fatal(err)
	}
}
