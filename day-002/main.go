package main

import "fmt"

/*
Write a function that takes a string of braces, and determines if the order of the braces is valid. It should return true if the string is valid, and false if it's invalid.

This Kata is similar to the Valid Parentheses Kata, but introduces new characters: brackets [], and curly braces {}. Thanks to @arnedag for the idea!

All input strings will be nonempty, and will only consist of parentheses, brackets and curly braces: ()[]{}.

What is considered Valid?
A string of braces is considered valid if all braces are matched with the correct brace.

Examples
"(){}[]"   =>  True
"([{}])"   =>  True
"(}"       =>  False
"[(])"     =>  False
"[({})](]" =>  False
*/

type Braces map[string]string

func NewBraces() Braces {
	return Braces{
		"(": ")",
		"[": "]",
		"{": "}",
	}
}

func main() {
	braces := NewBraces()
	testCases := "({()})){}[()]"
	valid := braces.ValidBraces(testCases)
	fmt.Println(valid)
}

func (b Braces) ValidBraces(str string) bool {
	var stack []string
	if len(str)%2 != 0 || len(str) == 0 {
		return false
	}
	for i := range str {
		if _, exists := b[str[i:i+1]]; exists {
			stack = append(stack, str[i:i+1])
		} else {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if b[last] != str[i:i+1] {
				return false
			}
		}
	}
	return len(stack) == 0
}
