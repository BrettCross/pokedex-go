package main

import "strings"

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	output := strings.Fields(lower)
	return output
}