package main

import (
  "fmt"
)

func main() {
  arr := []int{1,2,3,4,5}
  ans := getOtherProduct(arr)
  fmt.Println(ans)

  arr2 := []int{3,2,1}
  ans2 := getOtherProduct(arr2)
  fmt.Println(ans2)
}

func getOtherProduct(arr []int) []int {
  var ans []int
  var num int
  for _, a := range arr {
    num = 1
    for _, b := range arr {
      num = num * b
    }
    ans = append(ans, num/a)
  }
  return ans
}
