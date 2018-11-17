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
    go ref_tools.SleepInsert(&tempList, v, time.Duration(int64(v))*time.Millisecond)
  }
  time.Sleep(time.Duration(int64(vmax))*time.Millisecond)
  return tempList
}


func SpaghettiSort(list []float64) []float64 {
  tempList := []float64{}
  var allDone []bool
  increment := 0.1
  h := 0
  for i, v := range list {
    go ref_tools.SpaghettiInsert(&tempList, v, h)
  }
  for !tools.AllTrue(allDone) {
    pass
  }
  return tempList
}
