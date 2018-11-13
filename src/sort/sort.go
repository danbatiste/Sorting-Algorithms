package sort

import (
  //"fmt"
  "time"
)

type Generic interface{}

type GenericList []interface{}

func NewGenericList(values ...interface{}) GenericList {
  return values
}

func AllEqual(l1 []float64, l2 []float64) bool {
  if len(l1) != len(l2) {return false}
  for i:=0; i < len(l1); i++ {
    if l1[i] != l2[i] {return false}
  }
  return true
}

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

//  For a value E in a slice, waits E milliseconds and then inserts E into the list
//  O(n + k)
func SleepInsert(list *[]float64, value float64, t time.Duration) {
  time.Sleep(t)
  *list = append(*list, value)
}

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

//  Raises a line from 0 upwards. If a value is above the line,  it is added
//  to the slice.
func LineInsert(list *[]float64, value float64, h float64) bool {
  done := false
  increment = 0.1
  for h = 0; !done; h = h + increment {
    if value > h {
      *list = append(*list, value)
      done = true
    }
  }
  return done
}

func allTrue(list []bool) bool {
  init = true
  for _, e := range list {
    init = init && e
  }
  return init
}

func LineSort(list []float64) []float64 {
  tempList := []float64{}
  var allDone []bool{}
  for i, v := range list {
    allDone [i] = go LineInsert(&tempList, v, h)
  }
  for !allTrue(allDone) {
    ...
  }
  return tempList
}
