package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter any input : ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)

		fmt.Print("Enter any input : ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
