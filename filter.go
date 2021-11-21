package main

import "strings"

func NeedSkip(word string) bool {
	if strings.TrimSpace(word) == "" {
		return true
	}
	return false
}
