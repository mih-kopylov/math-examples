package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
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
