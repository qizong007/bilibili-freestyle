package main

import "strings"

var bullshit = []string{"的", "是", "好", "嗯", "后", "在", "于", "了"}
var specialChar = []string{" ", "+", "-", "*", "/", "[", "]", ".", ",", "。", "，"}

func FilterWords(words []string) []string {
	skippedList := make([]string, 0, len(words))
	for i := range words {
		if !needSkip(words[i]) {
			skippedList = append(skippedList, words[i])
		}
	}
	return removeListDuplication(skippedList)
}

func needSkip(word string) bool {
	word = strings.TrimSpace(word)
	if len(word) <= 3 {
		return true
	}
	if isSpecialChar(word) {
		return true
	}
	if isBullshit(word) {
		return true
	}
	return false
}

func isSpecialChar(word string) bool {
	return isStrInList(word, specialChar)
}

func isBullshit(word string) bool {
	for i := range bullshit {
		if strings.Contains(word, bullshit[i]) {
			return true
		}
	}
	return false
}

func isStrInList(word string, list []string) bool {
	for i := range list {
		if list[i] == word {
			return true
		}
	}
	return false
}

// 去重
func removeListDuplication(arr []string) []string {
	set := make(map[string]bool, len(arr))
	j := 0
	for _, v := range arr {
		if _, ok := set[v]; ok {
			continue
		}
		set[v] = true
		arr[j] = v
		j++
	}
	return arr[:j]
}
