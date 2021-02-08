package main

import (
	"bufio"
	"fmt"
)

const rules = `THE RULES OF NIM:
1. Nim is a mathematical game of strategy in which two players take turns removing (or "nimming") coins from a starting pile.
2. On each turn, the player removes either one, two, or three coins from the pile.
3. The player who removes the last coin from the pile wins the game!
`

func printRules(reader *bufio.Reader) {
	defer fmt.Println("Ok, let's play! We're going to play with 12 coins.\nI'll let you go first.")
	knowsRules := askUserYN("Do you know the rules of NIM?", reader)
	if !knowsRules {
		fmt.Println(rules)
	}
	return
}

func playNim(reader *bufio.Reader, coins int) {
	fmt.Printf("\nOk, your turn! there are %v coins", coins)
}
