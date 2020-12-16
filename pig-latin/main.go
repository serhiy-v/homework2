package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

)

func main() {
	var str string
	fmt.Printf("Enter your phrase:")
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	str = myscanner.Text()
	fmt.Printf("%s",transformPhrase(str))
}

func transformPhrase(str string) string{
	strSlice := strings.Fields(str)
	var result []string
	if len(strSlice) == 0{
		return "No word entered"
	}
	for _,v := range strSlice {
		v = transformToPigLatin((v))
		result = append(result, v)
	}
	return "PigLatin: " + strings.Join(result, " ")
}

func transformToPigLatin(str string)string{
	var ch rune
	for i, el := range str {
		if hasPunctuationMarks(el) {
			str = str[:i]
			ch = el
			return pigLatin(string(str)) + string(ch)
		}
	}
	return pigLatin(str)
}

func pigLatin(str string) string {
	strSlice := strings.Split(str, "")
	var suffix []string
	for i, v := range strSlice {
		if i == 0 && isVowel(v){
			return strings.Join(strSlice,"")+"yay"
		}
		if isVowel(v)  {
			suffix = strSlice[:i]
			strSlice = strSlice[i:]
			strSlice = append(strSlice,suffix...)
			return strings.Join(strSlice,"") + "ay"
		}
	}
	return str
}

func isVowel(ch string) bool{
	strSlice := strings.Split(ch,"")
	vowels  := []string {"A","a","E","e","I","i","O","o","U","u"}
	for _, ch := range strSlice{
		for _, v:=range vowels {
			if ch == v {
				return true
			}
		}
	}
	return false
}

func hasPunctuationMarks(char rune) bool{
	for _,el := range ".,!?()"{
		if char == el {
			return true
		}
	}
	return false
}
