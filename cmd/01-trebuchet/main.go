package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("./data/01-sample")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// create a slice to hold the list of numbers
	numbers := []int{}

	for scanner.Scan() {
		// keep the first and last numbers in memory
		// loop over each character
		// if it's a number, make it first and last
		// keep going until the end of the line
		// put the two numbers together
		// how best to do this?
		// ints -> characters -> concatenation -> back to int
		var firstNumChar, lastNumChar rune
		firstSet := false
		line := scanner.Text()
		reader := bufio.NewReader(strings.NewReader(line))
		for {
			char, _, err := reader.ReadRune()
			if err != nil {
				if err == io.EOF {
					break
				} else {
					log.Fatal(err)
				}
			}
			if unicode.IsDigit(char) {
				// fmt.Printf("the next char is a number: %q", char)
				if !firstSet {
					firstNumChar = char
					firstSet = true
				}
				lastNumChar = char
			}
			// fmt.Printf("%q", char)
		}
		numberString := fmt.Sprint(string(firstNumChar), string(lastNumChar))
		number, err := strconv.Atoi(numberString)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println(total)
}
