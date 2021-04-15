package main

func palindrome(x int) bool {

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0

	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x = x / 10
	}

	if x == revertedNumber || x == revertedNumber/10 {
		return true
	}

	return false
}

func isPalindromes(x int) bool {
	// always return false for negative numer
	if x < 0 {
		return false
	}
	// normal case: check if x eqauls to its reversed number
	reversed := 0
	for tmp := x; tmp != 0; tmp /= 10 {
		reversed = reversed*10 + tmp%10
	}
	return x == reversed
}
