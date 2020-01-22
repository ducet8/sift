package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode() & os.ModeNamedPipe == 0 {
		fmt.Println("Please pass data into sift")
	} else {
		choice, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		choice = choice - 1  //account for index starting at 0
		var line string
		var fields []string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line = scanner.Text()
			fields = strings.Fields(line)
			if choice <= len(fields) {
				fmt.Println(fields[choice])
			}
		}
	}
}