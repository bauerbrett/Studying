package main

import (
	"fmt"
)

//given a ransom note and a str, see if the letters in str can make ransom note

func ransomNote(note, str string) bool {
	noteChar := map[byte]int{}
	magChar := map[byte]int{}
	possible := true
	for i := 0; i < len(note); i++ {
		noteChar[note[i]]++
	}
	for i := 0; i < len(str); i++ {
		magChar[str[i]]++
	}
	for k, v := range noteChar {
		if magChar[k] < v {
			possible = false
		}
	}
	return possible
}
func main() {
	note := "apple"
	str := "pale"
	fmt.Println(ransomNote(note, str))
}
