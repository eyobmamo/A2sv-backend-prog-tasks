package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func fCoute(text string)map[string]int{
	text = strings.ToLower(text)

	cleanedText := strings.Map(func(r rune) rune{
		if unicode.IsLetter(r) || unicode.IsSpace(r){
			return r
		}
		return -1
	},text)
	fmt.Println(text)
	words := strings.Split(cleanedText, " ")
	count := make(map[string] int)
	for _,word := range words{
		if word != ""{
			count[word]++
		}
		
	}
	return count
}

func main() {
	reder := bufio.NewReader(os.Stdin)
	var text string
	fmt.Println("enter a text to count a word:")
	text, _ = reder.ReadString('\n')
	fmt.Println(fCoute(text))
}
