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

		if command == "exit"{
			break
		}else if strings.HasPrefix(command, "echo"){
			result := strings.TrimPrefix(command,"echo")
			fmt.Println(result)
		}else{
			fmt.Println(command+": command not found")
		}
		
		

		
	}
	
	
	
}
