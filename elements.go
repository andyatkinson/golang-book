package main

import "fmt"

func main() {
  elements := make(map[string]string)

  elements["H"] = "Hydrogen"
  elements["He"] = "Helium"
  elements["Li"] = "Lithium"
  elements["Be"] = "Beryllium"
  elements["B"] = "Boron"
  elements["C"] = "Carbon"
  elements["N"] = "Nitrogen"
  elements["O"] = "Oxygen"
  elements["F"] = "Fluorine"
  elements["Ne"] = "Neon"

  fmt.Println(elements["Li"])

  // accessing map values can return two elements
  // second value is whether the lookup was successful

  if name, ok := elements["Un"]; ok {
    fmt.Println(name, ok)
  }
}
