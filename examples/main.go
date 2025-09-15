package main

import (
	"fmt"

	"github.com/user/breeze"
)

func main() {
	fmt.Println("=== Breeze Hello World ===")

	// Simple AI response
	response := breeze.AI("Say hello in a creative way")
	fmt.Println("AI says:", response)

	// Clear conversation
	breeze.Clear()
	fmt.Println("\nConversation cleared!")
}
