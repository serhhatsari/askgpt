package utils

import "github.com/pterm/pterm"

func PrintDescription() {
	pterm.DefaultHeader.Println("Welcome to AskGPT!")
	pterm.Println(pterm.Blue("AskGPT is a CLI to interact with ChatGPT"))
	pterm.Println(pterm.Blue("You can ask any question and ChatGPT will answer it."))
	pterm.Println(pterm.Red("To exit, type \"exit\" or \"quit\" or press Ctrl C.\n"))
}
