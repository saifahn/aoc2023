package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	// store it in a map of map[int]string where int is the index of the number
	// then at the end, get the keys of the map, sort them
	// get the values of those keys from the map
	// put them together with concatenation
	// return the number
	fmt.Print(matchedNumbers)
	return 0
}

func main() {
	file, err := os.Open("data/01-sample")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		number := getFirstAndLastNumbersFromLine(line)

		// // read each character and check if it's a number
		// for {
		// 	char, _, err := reader.ReadRune()
		// 	if err != nil {
		// 		if err == io.EOF {
		// 			break
		// 		} else {
		// 			log.Fatal(err)
		// 		}
		// 	}
		// 	if unicode.IsDigit(char) {
		// 		if !isFirstSet {
		// 			firstNumChar = char
		// 			isFirstSet = true
		// 		}
		// 		lastNumChar = char
		// 		break
		// 	}
		// 	// naive implementation is to keep previous letter?
		// 	// nah, we should do like lookahead
		// 	// one
		// 	// two
		// 	// three
		// 	// four
		// 	// five
		// 	// six
		// 	// seven
		// 	// eight
		// 	// nine
		// 	// ten
		// }
		// // TODO: convert word numbers to regular numbers
		// numberString := fmt.Sprint(string(firstNumChar), string(lastNumChar))
		// number, err := strconv.Atoi(numberString)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		numbers = append(numbers, number)
	}

	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println(total)
}
