package main

import (
	"bufio"
	"fmt"
	myLogger "myapp/mylogger"
	"os"
	"time"
)

func main() {
	// read input from the user 5 times and write to a log
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)

	go myLogger.ListenForLog(ch)

	fmt.Println("Enter something")

	for i := 0; i < 5; i++ {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}
}
