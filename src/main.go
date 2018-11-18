package main

import (
  "fmt"
  "./sort"
  //"time"
)

type sortFunc
func Print(name string, func sortFunc, list []float64, on bool) {
  if on {
    fmt.Println(name)
    fmt.Println(sortfunc(list), "\n")
  }
}

func main() {
  list1 := []float64{1,0,10,13,11,100,51,3,4}
  //list2 := []float64{1}

  Print("BubbleSort", sort.BubbleSort, list1, true)

  Print("SleepSort", sort.SleepSort, list1, true)

  Print("ScanSort", sort.ScanSort, list1, false)

  Print("SpaghettiSort", sort.SpaghettiSort, list1, false)
}
