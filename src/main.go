package main

import (
  "fmt"
  "./sort"
  //"time"
)

func main() {
  list1 := []float64{1,10,13,11,100,51,3,4}

  fmt.Println("BubbleSort")
  a := sort.BubbleSort(list1)
  fmt.Println(a, "\n")

  fmt.Println("SleepSort")
  b := sort.SleepSort(list1)
  fmt.Println(b, "\n")
}
