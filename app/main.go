package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {	
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		command:=strings.TrimSpace(input)
		
		fmt.Println(command+": command not found")

		if command == "exit"{
			break
		}
	}
	
	
	
}
