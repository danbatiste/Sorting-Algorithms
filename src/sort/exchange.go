package sort

import (
  "./tools"
)


//  Typical bubble sort algorithm
func BubbleSort(list []float64) []float64 {
  var lastPass, currPass []float64
  tempList := list
  for !tools.AllEqual(lastPass, currPass) || currPass == nil {
    for i:=0; i + 1 < len(list); i++ {
        if tempList[i] > tempList[i+1] {
          tempList = tools.Swap(i, i+1, tempList)
        }
    }
    lastPass = currPass
    currPass = tempList
  }
  return tempList
}
