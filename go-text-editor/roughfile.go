package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	start()
}

func start() {
	// sample.txt
	mohfile := os.Args[1]

	data, err := os.ReadFile(mohfile)

	check_error(err)

	convertedstring := string(data)
	splittedString := splitSpace(convertedstring)

	fmt.Println(splittedString)

	// or
	// splitSpace(string(data))

	for position, word := range splittedString {
		if compare(word, "(hex)") == 1 {
			fmt.Println(position)
			fmt.Println("word is the same as (hex)")

		}
	}

}

func check_error(err error) {

	if err != nil {
		fmt.Printf("eror existing in the code")
		panic(err)
	}

}

func splitSpace(split string) []string {
	// split : Simply add          42
	string_to_split := strings.Fields(split)
	return string_to_split
}

func compare(asmau, xain string) int {
	if asmau == xain {
		return 1
	} else if asmau < xain {
		return 2
	} else if xain < asmau {
		return 3
	}
	return 4
}
