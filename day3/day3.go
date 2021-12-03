package main

import (
	"bufio"
	"log"
	"os"
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

	log.Println(input)

}
func gamma(input []string) int {
	zeroes := make([]int, 5, 5)
	ones := make([]int, 5, 5)

	for _, entry := range input {
		for i, bit := range entry {
			if bit == 48 {
				zeroes[i]++
			}
			if bit == 49 {
				ones[i]++
			}
		}
	}
}
