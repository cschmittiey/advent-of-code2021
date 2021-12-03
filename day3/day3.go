package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {

	input := make([]string, 0)

	file, err := os.Open("day3.demoinput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // for every line in the file
		input = append(input, scanner.Text())
	}
	if err != nil {
		log.Fatal(err)
	}

	zeroes := make([]int, 5)
	ones := make([]int, 5)

	for _, entry := range input {
		for i, bit := range entry {
			if bit == '0' {
				zeroes[i]++
			}
			if bit == '1' {
				ones[i]++
			}
		}
	}

	gammaResult, _ := strconv.ParseInt((string(gamma(zeroes, ones))), 2, 64)
	epsilonResult, _ := strconv.ParseInt((string(epsilon(zeroes, ones))), 2, 64)

	log.Println("Gamma is:", gammaResult)
	log.Println("Epsilon is:", epsilonResult)
	log.Println("Part 1 answer is:", gammaResult*epsilonResult)
	// log.Println("Epsilon is:", ^gammaResult)
	// really wanted to try and use bitwise negation on gamma to get epsilon
	// but I didn't quite figure it out before I got impatient and wanted to solve the puzzle.
	// but I mean also I guess I kind of did it in the epsilon function, it just wasn't as easy as doing ^gamma :(

}
func gamma(zeroes []int, ones []int) []byte {

	gamma := make([]byte, 5)

	for i := 0; i < 5; i++ {
		if zeroes[i] < ones[i] {
			gamma[i] = '1'
		} else {
			gamma[i] = '0'
		}
	}

	return gamma
}

func epsilon(zeroes []int, ones []int) []byte {
	epsilon := gamma(zeroes, ones) // hehe

	for i, x := range epsilon {
		if x == '1' {
			epsilon[i] = '0'
		}
		if x == '0' {
			epsilon[i] = '1'
		}
	} // this maybe wasn't the smartest way to do this. but i thought it was a little clever

	return epsilon
}
