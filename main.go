package main

import (
  "fmt"
  "./sort"
)


func main() {
  list := sort.NewGenericList(1,10,3,4)
  fmt.Println(sort.Swap(0,1,list))
  fmt.Println(sort.BubbleSort(list))
}
