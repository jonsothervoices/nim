package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, nim! Wanna play?")
	reader := bufio.NewReader(os.Stdin)
	str := readUserInput(reader)
	for !isYes(str) && !isNo(str) {
		fmt.Printf("Huh? what did you mean by %v?", str)
		str = readUserInput(reader)
	}
	if isNo(str) {
		fmt.Println("Ok, bye!")
		return
	}
	printRules(reader)
	playNim(reader, 12)
}
