package main

type stack []rune

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(str rune) {
	*s = append(*s, str)
}

// pop remove and return the top element of the stack
// return false when the stack is empty
func (s *stack) pop() (rune, bool) {
	if s.isEmpty() {
		return rune(0), false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}
