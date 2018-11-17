package sort

import (

)


//  Typical bubble sort algorithm
func BubbleSort(list []float64) []float64 {
  var lastPass, currPass []float64
  tempList := list
  for !AllEqual(lastPass, currPass) || currPass == nil {
    for i:=0; i + 1 < len(list); i++ {
        if tempList[i] > tempList[i+1] {
          tempList = Swap(i, i+1, tempList)
        }
    }
    lastPass = currPass
    currPass = tempList
  }
  return tempList
}
