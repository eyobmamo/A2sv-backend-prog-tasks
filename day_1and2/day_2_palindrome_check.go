package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func pali(text string) string{
	text = strings.ToLower(text)
	cleandedText := strings.Map(func(r rune) rune{
		if unicode.IsLetter(r){
			return r
		}
		return -1
	},text)
	n:=len(cleandedText)
	l:=0
	r := n-1
	// flag  := false
	for l < r{
		if (cleandedText[l] != cleandedText[r]){
			return "Not Palindrome"
		}
		l++
		r--

	}
	return "Palindrome"
	// if flag == true {
	// 	return fmt.Println(" Palindrome")
	// } else {
	// 	return fmt.Println("not palindrome")
	// }
	
}
func main(){
	var text string
	reder := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a text to cheak palindrome:")
	text ,_ =reder.ReadString('\n')
	fmt.Println(pali(text))
}