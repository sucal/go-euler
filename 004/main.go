package main

import "fmt"
import "math"

// A palindromic number reads the same both ways.
// The largest palindrome made from the product of
// two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product
// of two 3-digit numbers.

func pow10(x int) int {
	return int(math.Pow(float64(10), float64(x)))
}

func nthDigit(number int, n int) int {
	rem := number % pow10(n+1)
	return rem / pow10(n)
}

func numDigits(n int) int {
	num := n
	count := 1
	for ; (num / 10) != 0; count++ {
		num = num / 10
	}
	return count
}

func isPalindrome(n int) bool {

	nd := numDigits(n)
	k := nd - 1

	for i := 0; i < (nd / 2); i++ {

		ai := nthDigit(n, i)
		aki := nthDigit(n, k-i)

		if ai != aki {
			return false
		}
	}

	return true
}

func maxPalindromic(l int, u int) int {
	max := 0

	for i := l; i <= u; i++ {
		for j := l; j <= u; j++ {
			numToCheck := i * j
			if isPalindrome(numToCheck) && numToCheck > max {
				max = numToCheck
			}
		}
	}
	return max
}

func main() {

	fmt.Println("largest palindrome made from the product of two three digit numbers is ", maxPalindromic(100, 999))
}
