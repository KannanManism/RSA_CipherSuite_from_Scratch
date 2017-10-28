package main

import (
  "fmt"
  "math/rand"
  "strconv"
)

func main() {

  randomNumber := rand.Int31n(65536)
  fmt.Println("Random number is ", randomNumber)

  exponentiatedValue := squareAndMultiple(7,4,13)
  fmt.Println(exponentiatedValue)

}

func squareAndMultiple(a int, b int, c int) (int) {

  // FormatInt will provide the binary representation of a number
  binExp := strconv.FormatInt(int64(b),2)
  binExpLength := len(binExp)

  initialValue := a % c
  result := initialValue

  // Using the square and multipy algorithm to perform modular exponentation
  for i := 1; i < binExpLength; i++ {

    // 49 is the ASCII representation of 1 and 48 is the ASCII representation
    // of 0
    result = (result * result) % c
    if byte(binExp[i]) == byte(49) {
      result = (result * initialValue) % c
    }
  }
  return result

}
