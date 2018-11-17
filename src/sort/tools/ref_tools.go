package tools

import (
  "time"
)


//  Raises a line from 0 upwards. If a value is above the line,  it is added
//  to the slice.
func LineInsert(list *[]float64, value float64, h float64) bool {
  done := false
  increment := 0.1
  for h = 0; !done; h = h + increment {
    if value > h {
      *list = append(*list, value)
      done = true
    }
  }
  return done
}


//  For a value E in a slice, waits E milliseconds and then inserts E into the list
//  O(n + k)
func SleepInsert(list *[]float64, value float64, t time.Duration) {
  time.Sleep(t)
  *list = append(*list, value)
}
