package main

import (
  "fmt"
  "./sort"
)


func main() {
  list1 := sort.NewGenericList(1,10,13,11,100,3,4)

  fmt.Println("BubbleSort")
  fmt.Println(sort.BubbleSort(list1))

  fmt.Println("SleepSort")
  fmt.Println(sort.SleepSort(list1))
}
