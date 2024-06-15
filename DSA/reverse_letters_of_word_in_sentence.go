// Write an algorithm which reverses the letters of words in a sentence, like:
// Example input: I drive car

package main

import "fmt"

func main() {
	str := "I drive car"
	input := []byte(str)
	for i := 0; i < len(input); {
		start := i
		for i < len(input) && input[i] != ' ' {
			i += 1
		}
		swapReverse(input, start, i-1)
		i += 1
	}
	fmt.Println(string(input))
}

func swapReverse(input []byte, start, end int) {
	for start < end {
		temp := input[start]
		input[start] = input[end]
		start += 1
		input[end] = temp
		end -= 1
	}
}
