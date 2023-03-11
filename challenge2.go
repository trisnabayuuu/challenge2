package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("nilai i = %d\n", i)
	}
	j()
}

func j() {
	for j := 0; j < 11; j++ {
		if j == 5 {
			for _, c := range []rune{'\u0421', '\u0410', '\u0428', '\u0410', '\u0420', '\u0412', '\u041e'} {
				fmt.Printf("%c \n", c)
			}
		} else {
			fmt.Printf("nilai J = %d\n", j)
		}
	}
}
