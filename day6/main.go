package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	feesh := make([]uint8, 0)

	file, err := os.Open("day6.input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	initialInput := strings.Split(scanner.Text(), ",")
	if err != nil {
		log.Fatal(err)
	}

	for _, chunk := range initialInput {
		fish, _ := strconv.ParseUint(chunk, 10, 8)
		fish8 := uint8(fish)
		feesh = append(feesh, fish8)

	}

	for i := 0; i < 256; i++ {
		log.Println(i)

		for i, fish := range feesh {
			if fish == 0 {
				feesh = append(feesh, 8)
				feesh[i] = 6
			}
			if fish > 0 {
				feesh[i] = fish - 1
			}

		}
		log.Println(i)

	}

	log.Println(len(feesh))

}
