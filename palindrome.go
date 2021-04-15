package main

import "strconv"

func palindrome(x int) {

	str := strconv.Itoa(x)

	for c := range str {

		println(string(rune(c)))
	}

}
