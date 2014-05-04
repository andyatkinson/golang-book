package main

import "fmt"

// recursion example
func factorial(x uint) uint {
  if x == 0 {
    return 1
  }
  return x * factorial(x-1)
}

func main() {
  fmt.Println(factorial(5))
}
