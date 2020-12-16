package main

import (
	"bufio"
	"fmt"
	"os"
)

var letters = map[rune]rune{
	'a': '1',
	'A': '1',
	'e': '2',
	'E': '2',
	'i': '3',
	'I': '3',
	'o': '4',
	'O': '4',
	'u': '5',
	'U': '5',
}

var numbers = map[rune]rune{
	'1': 'a',
	'2': 'e',
	'3': 'i',
	'4': 'o',
	'5': 'u',
}



func main() {
	var str string
	fmt.Printf("Enter your phrase:")
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	str = myscanner.Text()
	if len(str) == 0{
		fmt.Printf("No word entered")
	}else {
		fmt.Printf("Encoded: %s\n", encode(str))
		fmt.Printf("Decoded: %s", encode(str))
	}
}

func encode(str string) string{
	result := ""
	for _, el := range str{
		letter, ok := letters[el]
		if ok {
			result += string(letter)
		}else {
			result += string(el)
		}
	}
	return result
}

func decode(str string) string{
	result := ""
	for _, el := range str{
		number, ok := numbers[el]
		if ok {
			result += string(number)
		}else {
			result += string(el)
		}
	}
	return result
}