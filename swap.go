package main

import "fmt"

func swap(x *int, y *int) {
  tmp := *x // dereference pointer value
  *x = *y // swap pointer values
  *y = tmp // assign y pointer to tmp value
}

func main() {
  x := 1
  y := 2
  fmt.Println("x ",x, "y ",y)
  swap(&x, &y)
  fmt.Println("x ",x, "y ",y)
}
