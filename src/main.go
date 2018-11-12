package main

import (
  "fmt"
  "./sort"
  "time"
)


func main() {
  list1 := []float64{1,10,13,11,100,3,4}

  fmt.Println("BubbleSort")
  fmt.Println(sort.BubbleSort(list1))

  fmt.Println("SleepSort")
  a := sort.SleepSort(list1)
  fmt.Println(a)
  time.Sleep(100000)
  fmt.Println(a)

}
