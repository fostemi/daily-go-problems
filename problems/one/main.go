package main

import (
  "fmt"
)

func main() {
  nums := []int{10, 15, 4, 7}
  k := 17
  res := canSumTo(nums, k)
  fmt.Println(res)
}

func canSumTo(nums []int, k int) bool {
  for _, n1 := range nums {
    for _, n2 := range nums {
      if (n1 + n2) == k {
        return true
      }
    }
  }
  return false
}
