package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello, nim! Wanna play online?")
	reader := bufio.NewReader(os.Stdin)
	str := readUserInput(reader)
	for !isYes(str) && !isNo(str) {
		fmt.Printf("Huh? what did you mean by %v?", str)
		str = readUserInput(reader)
	}
	if isNo(str) {
		printRules(reader)
		playNim(reader, 12)
		return
	}
	//online stuff
	fmt.Println("Go to http://localhost:8080/ to play--I'll see you there!")
	http.Handle("/", newAPI())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
