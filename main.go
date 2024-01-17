package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	err := runApplication()
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}
	waitForEnter()
}

func waitForEnter() {
	fmt.Println("Нажмите Enter")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}
