package main

import (
	"fmt"
	"unicode/utf8"
)

func PrintRune(ch rune){
	fmt.Printf("%#U \n", ch)
}

func main(){
	const s = "abc"
	for ix := range s {
		fmt.Println(s[ix])
	}

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	for ix, val := range s {
		fmt.Printf("%d %#U\n", ix, val)
	}


	// print rune
	fmt.Println("Print Runes")
	for _, val := range s {
		PrintRune(val)
	}
}