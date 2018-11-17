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

//  TODO: Implement for negative numbers
func SpaghettiSort(list []float64) []float64 {
  tempList := []float64{}
  var allDone []bool
  var h float64
  increment := 0.1
  h = 0
  for i, v := range list {
    go tools.SpaghettiInsert(&tempList, &allDone[i], v, &h)
  }
  for !tools.AllTrue(allDone) {
    h += increment
  }
  return tempList
}
