package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

var stringsToDigits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

var numStrings = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

// create a function to check
func getFirstAndLastNumbersFromLine(currentLine string) int {
	matchedNumbers := map[int]string{}
	// for each number string and char
	// check for matches in the line
	for _, s := range numStrings {
		re := regexp.MustCompile(s)
		matches := re.FindAllStringIndex(currentLine, -1)
		if matches == nil {
			continue
		}
		for _, match := range matches {
			matchedNumbers[match[0]] = s
		}
	}
	// then at the end, get the keys of the map, sort them
	indexKeys := make([]int, len(matchedNumbers))
	i := 0
	for k := range matchedNumbers {
		indexKeys[i] = k
		i++
	}
	slices.Sort(indexKeys)
	// get the values of those keys from the map
	firstDigit := stringsToDigits[matchedNumbers[indexKeys[0]]]
	secondDigit := stringsToDigits[matchedNumbers[indexKeys[len(indexKeys)-1]]]

	// put them together with concatenation
	twoDigitString := fmt.Sprint(string(firstDigit), string(secondDigit))
	twoDigitNumber, err := strconv.Atoi(twoDigitString)
	if err != nil {
		log.Fatal(err)
	}

	// return the number
	return twoDigitNumber
}

func main() {
	file, err := os.Open("data/01-pt2-sample")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		number := getFirstAndLastNumbersFromLine(line)
		numbers = append(numbers, number)
	}

	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println(total)
}
