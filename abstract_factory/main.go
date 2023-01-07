package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	WINDOWS = "windows"
	MAC     = "mac"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var system string
	fmt.Printf("Which system are you in? (%s, %s)\n", WINDOWS, MAC)
	if scanner.Scan() {
		system = scanner.Text()
	}

	var factory GUIFactory
	switch system {
		case WINDOWS:
			factory = WinFactory{}
		case MAC:
			factory = MacFactory{}
		default:
			fmt.Println("System not supported")
	}

	app := NewApplication(factory)
	app.CreateUI()
}
