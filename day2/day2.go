package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	magnitude int
}

func main() {

	file, err := os.Open("day2.input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	instructions := make([]instruction, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // for every line in the file
		instructionText := scanner.Text()
		instructionSplit := strings.Split(instructionText, " ")
		direction := instructionSplit[0]
		magnitude, _ := strconv.Atoi(instructionSplit[1])

		instructions = append(instructions, instruction{direction, magnitude})
	}
	if err != nil {
		log.Fatal(err)
	}

	// parse instructions
	horizontal := 0
	vertical := 0

	for _, currentInstruction := range instructions {
		if currentInstruction.direction == "forward" {
			horizontal = horizontal + currentInstruction.magnitude
		}
		if currentInstruction.direction == "up" {
			vertical = vertical + currentInstruction.magnitude
		}
		if currentInstruction.direction == "down" {
			vertical = vertical + currentInstruction.magnitude
		}
	}

	log.Println(horizontal)
	log.Println(vertical)

}
