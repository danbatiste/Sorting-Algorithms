package tools

import (

)

func Min(list []float64) float64 {
  min := 0.0
  for _, e := range list {
    if e < min {min = e}
  }
  return min
}

func Max(list []float64) []float64 {
  max := 0.0
  for _, e := range list {
    if e > max {max = e}
  }
  return max
}

func AllTrue(list []bool) bool {
  init := true
  for _, e := range list {
    init = init && e
  }
  return init
}

func AllEqual(l1 []float64, l2 []float64) bool {
  if len(l1) != len(l2) {return false}
  for i:=0; i < len(l1); i++ {
    if l1[i] != l2[i] {return false}
  }
  return true
}

//  Swaps two elements in a slice
func Swap(index1 int, index2 int, list []float64) []float64 {
  var tempList []float64
  for i:=0; i < len(list); i++ {
    switch {
    case i == index1:
      tempList = append(tempList, list[index2])
    case i == index2:
      tempList = append(tempList, list[index1])
    default:
      tempList = append(tempList, list[i])
    }
  }
  return tempList
}
