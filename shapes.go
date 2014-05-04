package main

import ("fmt"; "math")

// struct example
type Circle struct {
  x, y, r float64
}

// another struct
type Person struct {
  Name string
}
func (p *Person) Talk() {
  fmt.Println("Hi, my name is", p.Name)
}
// anonymous fields/embedded type
type Android struct {
  Person
}

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
  l := distance(x1, y1, x1, y2)
  w := distance(x1, y1, x2, y1)
  return l * w
}

// shows a receiver, name and type
func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

func main() {
  var rx1, ry1 float64 = 0, 0
  var rx2, ry2 float64 = 10, 10
  circle := Circle{0, 0, 5}

  fmt.Println("rectangle area", rectangleArea(rx1, ry1, rx2, ry2))
  fmt.Println("circle area", circle.area())
  p := Person{Name: "Bob"}
  p.Talk()
  a := Android{Person{Name: "Android"}} // more direct way to assign parent class property?
  a.Talk()
}
