package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open up the file!

	file, err := os.Open("day1.input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// make the slice to put numbers in
	numbers := make([]int, 0)

	// read the numbers into the slice
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // for every line in the file
		newNumber, err := strconv.Atoi(scanner.Text()) // grab a line, convert the ascii to an int
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, newNumber)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// let's actually do the work here:

	counter := 0

	for i, number := range numbers { // reeeee
		if i > 0 {
			if number > numbers[i-1] {
				counter++
			}
		}
	}

	log.Println("=== Done! The answer is: ", counter)

}
