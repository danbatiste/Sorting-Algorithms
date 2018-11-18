package main

import (
  "fmt"
  "./sort"
  //"time"
)

func main() {
  list1 := []float64{1,0,10,13,11,100,51,3,4}
  //list2 := []float64{1}

  fmt.Println("BubbleSort")
  a := sort.BubbleSort(list1)
  fmt.Println(a, "\n")

  fmt.Println("SleepSort")
  b := sort.SleepSort(list1)
  fmt.Println(b, "\n")

  fmt.Println("SpaghettiSort")
  c := sort.SpaghettiSort(list1)
  fmt.Println(c, "\n")
}
