package main

import "fmt"
import "math"

// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143 ?

func largestIntLessThanSqrt(n int) int {
  return int(math.Ceil(math.Sqrt( float64(n) )))
}

func isPrime(n int) bool {

  upper := largestIntLessThanSqrt(n)

  for i:= 2; i < upper; i++ {
    if n % i == 0 {
      return false
    }
  }
  return true
}

func largestPrimeFactor(n int) int {
  upper := largestIntLessThanSqrt(n)
  result := 0

  for i:= 2; i <= upper; i++ {
    if (n % i == 0) && isPrime(i) {
      result = i
    }
  }

  return result
}

func main()  {
  fmt.Println(largestPrimeFactor(600851475143))
}
