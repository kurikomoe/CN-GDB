package main

import "fmt"

var myst = "Shall we?"

func main () {
  fmt.Println ("Before assignment")
  st := "Hello, world!" // this intentionally shadows the global "st"
  fmt.Println (st) // set breakpoint 1 here
  fmt.Println (myst)
}
