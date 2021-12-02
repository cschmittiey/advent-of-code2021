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

	file, err := os.Open("day2.demoinput.txt")
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

	horizontal, vertical := parseInstructionsPartOne(instructions, horizontal, vertical)

	log.Println("Horizontal Position:", horizontal)
	log.Println("Vertical Position:", vertical)
	log.Println("Answer for Day 2, Part 1:", horizontal*vertical)

	horizontal2, vertical2, aim := parseInstructionsPartTwo(instructions, horizontal, vertical)

	log.Println("Horizontal Position:", horizontal2)
	log.Println("Vertical Position:", vertical2)
	log.Println("Aim:", aim)
	log.Println("Answer for Day 2, Part 1:", horizontal2*vertical2)

}

func parseInstructionsPartOne(instructions []instruction) (int, int) {
	horizontal := 0
	vertical := 0

	for _, currentInstruction := range instructions {
		if currentInstruction.direction == "forward" {
			horizontal = horizontal + currentInstruction.magnitude
		}
		if currentInstruction.direction == "up" {
			vertical = vertical - currentInstruction.magnitude
		}
		if currentInstruction.direction == "down" {
			vertical = vertical + currentInstruction.magnitude
		}
	}
	return horizontal, vertical
}

func parseInstructionsPartTwo(instructions []instruction) (int, int, int) {
	horizontal := 0
	vertical := 0
	aim := 0

	for _, currentInstruction := range instructions {
		if currentInstruction.direction == "forward" {
			horizontal = horizontal + currentInstruction.magnitude
			vertical = vertical + (currentInstruction.magnitude * aim)
		}
		if currentInstruction.direction == "up" {
			aim = aim - currentInstruction.magnitude
		}
		if currentInstruction.direction == "down" {
			aim = aim + currentInstruction.magnitude
		}
	}
	return horizontal, vertical, aim
}
