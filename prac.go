/*
package main

import ("fmt"
	   "strconv"
)

func main(){
fmt.Println(HexToDecimal("1E"))
}

func HexToDecimal(hexString string) (int64, error) {
	converted, err := strconv.ParseInt(hexString, 16, 64)
	return converted, err
}
*/

/*
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(BintoDec("10"))
}

func BintoDec(hexString string) (int64, error) {
	con, err := strconv.ParseInt(hexString, 2, 64)
	return con, err
}
*/

//Toupper

/*
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(toUpper("hElLo"))
}
func toUpper(word string) string {
	return string(unicode.ToUpper(rune(word[0]))) + strings.ToLower(word[1:])
}
*/

/*package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
func main() {
	fmt.Println(ToLower("HELLO"))
}

func ToLower(word string) string {
	return string(unicode.ToLower(rune(word[0]))) + strings.ToUpper(word[1:])
}

*/

/*
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ToLower("HELLO"))
}
func ToLower(word string) string {
	first := strings.ToLower((string(word[0])))
	rest := strings.ToUpper(word[1:])
	return first + rest
}
*/

/*
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ToLower("HELLO"))
}

func ToLower(word string) string {
	first := strings.ToLower(string(word[0]))
	rest := strings.ToUpper(word[1:])
	return first + rest
}

*/

/*
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toUpper("hELLO"))
}

func toUpper(word string) string {
	first := strings.ToUpper(string(word[0]))
	last := strings.ToLower(word[1:])
	return first + last
}

*/

//LOWER AND UPPER CASE STILL

/*
package main // Every Go program starts with a package. "main" means this program can run.

import (
	"fmt"     // Used to print things to the terminal
	"strings" // Used for working with text (uppercase, lowercase, etc.)
)

func main() {

	// Call the changeWord function and print the result
	fmt.Println(changeWord("HELLO"))

}

// changeWord takes a string and returns a modified string
func changeWord(word string) string {

	// If the word is empty, return an empty string
	// This prevents the program from crashing when accessing word[0]
	if len(word) == 0 {
		return ""
	}

	// Take the first letter of the word
	// word[0] gets the first character but it is a byte
	// string(...) converts it back to a string
	firstLetter := string(word[0])

	// Convert the first letter to lowercase
	firstLetter = strings.ToLower(firstLetter)

	// Get the rest of the word starting from position 1
	restOfWord := word[1:]

	// Convert the rest of the word to uppercase
	restOfWord = strings.ToUpper(restOfWord)

	// Combine the lowercase first letter with the uppercase rest
	result := firstLetter + restOfWord

	// Return the final word
	return result
}

*/

/*

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(changeWord("Hello"))

}

func changeWord(word string) string {

	if len(word) == 0 {
		return ""
	}

	runes := []rune(word)

	first := strings.ToLower(string(runes[0]))
	rest := strings.ToUpper(string(runes[1:]))

	return first + rest
}
*/

// FOR TITLE CASE STILL

/*

package main

import (
	"fmt"
	"unicode"
)

func main() {

	fmt.Println(titleCase("hEllo"))
}

func titleCase(s string) string {
	// nothing to do with empty strings

	if s == "" {
		return s
	}
	//convert the whole string to a rune for safe unicode
	runes := []rune(s)
	//uppercae the very first character
	runes[0] = unicode.ToUpper(runes[0])
	// lowecase every character after the first
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)

}
*/

/*
// TitleCase

package main

import (
	"fmt"
	"unicode"
)

func main() {

	// Calling the titleCase function
	// The function will format the word
	result := titleCase("hEllo")

	// Print the result to the terminal
	fmt.Println(result)

}

// titleCase converts a word so that
// the first letter is uppercase
// and the rest are lowercase
func titleCase(s string) string {

	// If the string is empty, return it immediately
	if s == "" {
		return s
	}

	// Convert the string into a slice of runes
	// This allows safe handling of Unicode characters
	runes := []rune(s)

	// Convert the first character to uppercase
	runes[0] = unicode.ToUpper(runes[0])

	// Convert every character after the first to lowercase
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}

	// Conv

*/

///Title case to fix an entire sentence

/*
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(titleSentence("hEllo wORLD frOM go"))
}

func titleSentence(s string) string {

	words := strings.Split(s, " ")

	for i := 0; i < len(words); i++ {
		words[i] = titleCase(words[i])
	}

	return strings.Join(words, " ")
}

func titleCase(s string) string {

	if s == "" {
		return s
	}

	runes := []rune(s)

	runes[0] = unicode.ToUpper(runes[0])

	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}

	return string(runes)
}

*/

///Title case to fix an entire sentence 2

package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "hello world"

	result := strings.Title(text)

	fmt.Println(result)

}

/*
//shortest version of titleCase

package main

import (
	"fmt"
)
*/

func main() {
	fmt.Println(ToCap("hEllo"))
}

func capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

