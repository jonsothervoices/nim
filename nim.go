package main

import (
	"bufio"
	"fmt"
)

type game struct {
	Coins int
	Move  int
	State string
}

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
	fmt.Printf("Ok, your turn! There are %v coins\n", coins)
	fmt.Println("How many coins do you want to remove?")
	//Read user input
	i := readValidInt(reader, 1, 3)
	//take away that many coins
	coins -= i

	fmt.Printf("Now there are %v coins\n", coins)
	//take away 4-the previous user input
	fmt.Printf("My turn! So, I'll take %v coins\n", 4-i)
	coins -= (4 - i)
	//check total coin number==0
	if coins == 0 {
		//if we're done,
		fmt.Printf("It looks like I win!\nThanks for playing, maybe you'll beat me next time!\n")
		if askUserYN("Do you want to play again?", reader) {
			playNim(reader, 12)
			return
		}
		fmt.Println("Ok, bye!")
		return
	}
	//if not, recurse
	playNim(reader, coins)
}

func (daGame *game) play() error {
	if daGame.Move < 1 || daGame.Move > 3 {
		return fmt.Errorf("Invalid number: must be between 1 and 3")
	}
	daGame.Coins -= daGame.Move
	daGame.Move = 4 - daGame.Move
	daGame.Coins -= daGame.Move
	if daGame.Coins == 0 {
		daGame.State = "you lose!"
	}
	return nil
}
