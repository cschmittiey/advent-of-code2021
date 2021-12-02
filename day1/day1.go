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
	threesomes := make([]int, 0)

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

	for i, number := range numbers {
		if i > 0 && i < (len(numbers)-1) {
			threesome := number + numbers[i-1] + numbers[i+1]
			threesomes = append(threesomes, threesome)
		}
	}

	// reeeee
	partOneAnswer := countIncreases(numbers)
	partTwoAnswer := countIncreases(threesomes)

	log.Println("=== Done! Part One's answer is: ", partOneAnswer)
	log.Println("=== Done! Part Two's answer is: ", partTwoAnswer)

}

func countIncreases(numbers []int) int {
	numberOfIncreases := 0

	for i, number := range numbers {
		if i > 0 {
			if number > numbers[i-1] {
				numberOfIncreases++
			}
		}
	}
	return numberOfIncreases
}
