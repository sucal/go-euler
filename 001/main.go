package main

import "fmt"
import "flag"

// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below 1000.

type Predicate func(n int) bool

func sumBelow(n int, p Predicate) int {
  sum := 0
  for i:=0; i < n; i++ {
    if p(i) == true {
      sum += i
    }
  }
  return sum
}

func main()  {
  limit := flag.Int("n", 10, "the upper limit of natural numbers")
  flag.Parse()

  result := sumBelow(*limit, func(n int) bool {
      if n%3 == 0 || n%5 == 0 {
        return true
      } else {
        return false
      }
  })

  fmt.Println(result)
}
