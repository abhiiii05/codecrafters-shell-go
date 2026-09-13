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
			prefix := strings.TrimPrefix(command,"echo")
			result := strings.TrimSpace(prefix)
			fmt.Println(result)
		}else if strings.HasPrefix(command,"type"){
			prefix := strings.TrimPrefix(command," type")
			fmt.Println("here")
			if (prefix == "echo" || prefix =="exit" || prefix=="type"){
				fmt.Println(prefix + "is a shell builtin")
			}
		}else{
			fmt.Println(command+": command not found")
		}
		
		

		
	}
	
	
	
}
