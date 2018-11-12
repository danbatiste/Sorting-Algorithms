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

func AllEqual(l1 []interface{}, l2 []interface{}) bool {
  if len(l1) != len(l2) {return false}
  for i:=0; i < len(l1); i++ {
    if l1[i] != l2[i] {return false}
  }
  return true
}

func IsGreaterThan(v1, v2 interface{}) bool {
  switch t1 := v1.(type) {
  case int:
    t2, _:= v2.(int)
    return t1 > t2
  default:
    return false
  }
}

func Swap(index1 int, index2 int, list []interface{}) []interface{} {
  var tempList []interface{}
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

func BubbleSort(list []interface{}) []interface{} {
  var lastPass, currPass []interface{}
  tempList := list
  for !AllEqual(lastPass, currPass) || currPass == nil {
    for i:=0; i + 1 < len(list); i++ {
        if IsGreaterThan(tempList[i], tempList[i+1]) {
          tempList = Swap(i, i+1, tempList)
        }
    }
    lastPass = currPass
    currPass = tempList
  }
  return tempList
}

func SleepInsert(list *[]float64, value float64, t time.Duration) {
  time.Sleep(t)
  *list = append(*list, value)
}

func SleepSort(list []float64) []float64 {
  var tempList []float64
  for _, v := range list {
    go SleepInsert(&tempList, v, time.Duration(int64(v*1000)))
  }
  return tempList
}
