package main

import "log"

func main() {
	log.Println(isValid("()"))
}

var mappings map[rune]rune

func initMappings() {
	mappings = make(map[rune]rune)
	mappings[')'] = '('
	mappings['}'] = '{'
	mappings[']'] = '['
}

func isValid(s string) bool {
	initMappings()

	var stackOfRune stack
	for _, char := range s {
		// if the character is a closing bracket
		if mappings[char] != rune(0) {
			topElement, valid := stackOfRune.pop()
			// when the mapping of the opening bracket and the top element doesn't match
			// or if we have a closing bracket and the stack is empty
			// we have a wrong expression
			if mappings[char] != topElement || !valid {
				return false
			}
		} else {
			// when we have an opening bracket, simply push to the stack
			stackOfRune.push(char)
		}
	}
	return stackOfRune.isEmpty()
}
