package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/user/breeze"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: breeze <prompt>")
		fmt.Println("       breeze chat <prompt>")
		fmt.Println("       breeze code <prompt>")
		fmt.Println("       breeze clear")
		return
	}

	switch args[0] {
	case "chat":
		if len(args) < 2 {
			fmt.Println("Usage: breeze chat <prompt>")
			return
		}
		response := breeze.Chat(strings.Join(args[1:], " "))
		fmt.Println(response)
	case "code":
		if len(args) < 2 {
			fmt.Println("Usage: breeze code <prompt>")
			return
		}
		response := breeze.Code(strings.Join(args[1:], " "))
		fmt.Println(response)
	case "clear":
		breeze.Clear()
		fmt.Println("Conversation cleared.")
	default:
		// Default to ai
		prompt := strings.Join(args, " ")
		response := breeze.AI(prompt)
		fmt.Println(response)
	}
}
