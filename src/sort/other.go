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
    go SleepInsert(&tempList, v, time.Duration(int64(v))*time.Millisecond)
  }
  time.Sleep(time.Duration(int64(vmax))*time.Millisecond)
  return tempList
}


func LineSort(list []float64) []float64 {
  tempList := []float64{}
  var allDone []bool
  for i, v := range list {
    allDone [i] = go LineInsert(&tempList, v, h)
  }
  for !allTrue(allDone) {
    ...
  }
  return tempList
}
