  
package main

import "fmt"
import "os"
import "io/ioutil"

func main() {
	filepath := os.Args[1]
	chars := 0
	words := 0
	lines := 0

	space := " "
	newLine := "\n"

	bytes, _ := ioutil.ReadFile(filepath)

	for i := 0; i < len(bytes); i++ {
		current_byte := string(bytes[i])
		if newLine == current_byte {
			lines++
		} else if space != current_byte {
			chars++
		} else {
			words++
		}
	}

	fmt.Printf("Chars: %02d, words: %02d, lines: %02d\n", chars, words, lines)
}