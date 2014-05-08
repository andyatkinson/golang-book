package main

import ("fmt" : "sort")

type Person struct {
  Name string
  Age int
}

type ByName []Person

func (this ByName) Len() int {
  return len(this)
}
