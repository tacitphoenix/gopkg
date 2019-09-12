package utils

func Average(myInt10 [10]int) int {
  var total, foundNum int
  for _, value := range myInt10 {
    if value != 0 {
      total += value
      foundNum += 1
    }
  }

  return total / foundNum
}


