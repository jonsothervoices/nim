package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readUserInput(reader *bufio.Reader) string {
	str, err := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	if err != nil {
		//2 ways to write to stderr
		fmt.Fprintln(os.Stderr, "Input read error: ", err)
		//os.Stderr.WriteString(fmt.Sprintf("Input read error: %v\n", err))
		panic(err)
	}
	return str
}

func askUserYN(q string, reader *bufio.Reader) bool {
	fmt.Println(q)
	ans := readUserInput(reader)
	for !isYes(ans) && !isNo(ans) {
		fmt.Printf("Huh? What did you mean by %v?", ans)
		ans = readUserInput(reader)
	}
	return isYes(ans)
}

func isYes(s string) bool {
	mYes := map[string]bool{"y": true, "yes": true, "sure": true, "yeah": true, "ok": true, "okay": true, "yup": true, "why not": true}
	_, ok := mYes[strings.ToLower(s)]
	return ok
}

func isNo(s string) bool {
	mNo := map[string]bool{"n": true, "no": true, "no thanks": true, "no thank you": true, "maybe next time": true, "nah": true, "nah bruh": true, "next time": true}
	_, ok := mNo[strings.ToLower(s)]
	return ok
}

func readValidInt(reader *bufio.Reader, l, u int) int {
	s := readUserInput(reader)
	ret, err := strconv.Atoi(s)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		fmt.Printf("\nThat's not how this works...that's not how any of this works\nLet's try this again...\nHow many coins would you like to remove?\n")
		return readValidInt(reader, l, u)
	}
	if ret < l || ret > u {
		fmt.Printf("%v is not a valid number! Input a number between %v and %v\nLet's try this again...\nHow many coins would you like to remove?\n", ret, l, u)
		return readValidInt(reader, l, u)
	}
	return ret
}
