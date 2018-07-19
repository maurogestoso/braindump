package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	entry := strings.Join(os.Args[1:], " ")

	journal, err := os.OpenFile("braindump.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer journal.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = journal.WriteString(fmt.Sprintf("%s\n", entry))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
