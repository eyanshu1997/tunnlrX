package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tunnlrx <command>")
		return
	}
	switch os.Args[1] {
	case "up":
		fmt.Println("🔗 Starting tunnel on port 8080 (placeholder)")
	default:
		fmt.Println("Unknown command")
	}
}
