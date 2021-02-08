package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Printf("Huh? what did you mean by %v?", ans)
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
