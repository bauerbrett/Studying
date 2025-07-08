package main

import (
	"fmt"
	"math"
)

func maxBalloon(s string) int {
	balloonMap := map[rune]int{}
	balloon := "balloon"
	if len(s) < 7 {
		return 0
	}
	for _, char := range balloon {
		balloonMap[char] = 0
	}
	for _, char := range s {
		balloonMap[char]++
	}
	small := math.MaxInt32
	for _, num := range balloonMap {
		if num < small && num > 0 {
			small = num
		}
	}
	if small == math.MaxInt32 {
		return -1
	}
	return small

}

func main() {
	str := "balloonballoon"
	str1 := "bbaall"
	fmt.Println(maxBalloon(str), maxBalloon(str1))
}
