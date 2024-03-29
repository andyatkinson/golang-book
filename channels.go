package main

import (
  "fmt"
  "time"
)

// blocks until printer is ready to receive the msg
func pinger(c chan string) {
  for i := 0; ; i++ {
    c <- "ping"
  }
}

func printer(c chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func ponger(c chan string) {
  for i := 0; ; i++ {
    c <- "pong"
  }
}

func main() {
  var c chan string = make(chan string)

  go pinger(c)
  go printer(c)
  go ponger(c)

  var input string
  fmt.Scanln(&input)
}
