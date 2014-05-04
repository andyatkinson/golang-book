package main

import "fmt"

// closure example, "x" and "increment" form the closure
func main() {
  x := 0
  increment := func() int {
    x++
    return x
  }
  fmt.Println(increment())
  fmt.Println(increment())
}
