package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./data/01-sample")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
}
