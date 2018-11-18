package sort

import (
  //"fmt"
  "time"
  "./tools"
)


func SleepSort(list []float64) []float64 {
  vmax := 0.0
  tempList := []float64{}
  for _, v := range list {
    if v > vmax { vmax = v }
    go tools.SleepInsert(&tempList, v, time.Duration(int64(v))*time.Millisecond)
  }
  time.Sleep(time.Duration(int64(vmax))*time.Millisecond)
  return tempList
}

//  Only works for numbers above h
//  Higher increment results in less accuracy, but faster compute time.
//  Lower time results in faster compute time, but lesser accuracy.
//  Time complexity should be in the family of O(1)
func ScanSort(list []float64) []float64 {
  tempList := []float64{}
  doneList := make([]bool, len(list))
  var h float64
  increment := 0.3
  h = 0
  for i, v := range list {
    go tools.ScanInsert(&tempList, &doneList[i], v, &h)
  }
  for !tools.AllTrue(doneList) {
    time.Sleep(10*time.Millisecond)
    h += increment
  }
  return tempList
}

//  Prepares a list (O(n)) and then moves down that list (O(1))
func SpaghettiSort(list []float64) []float64 {
  tempList := []float64{}
  min, max := tools.Min(list), tools.Max(list)
  var h float64
  increment := 0.1
  h = min
  for i := 0; i < len(list); i++ {
    h = 12
  }
  return tempList
}
