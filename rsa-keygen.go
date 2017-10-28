package main

import (
  "fmt"
  //"math/rand"
  "strconv"
)

func main() {

  //randomNumber := rand.Int31n(65536)
  randomNumber := 281
  fmt.Println("Random number is ", randomNumber)

  exponentiatedValue := squareAndMultiple(7,4,13)
  fmt.Println(exponentiatedValue)

  // Check for a prime number
  // I'm hardcoding the value of K in primality test to 5
  accuracyFactor := 5;
  isaPrimeNumber(randomNumber,accuracyFactor)

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

func isaPrimeNumber(number int, accuracyFactor int) {

  // First finding the value of r, d as per equation ;
  // d * 2pow(r) = n - 1

  varNumber := float64(number - 1)
  r := 2
  // exponentitalR is 2powr(r)
  exponentitalR := float64(2)

  for true {
    x := varNumber/exponentitalR
    fmt.Println(x)
    if (x == float64(int64(x))) {
    // Fixing value 10000000000 for calculation purpose
    // To resue the squareAndMultiple algorithm but not affect the modulo part
      r = r + 1
      exponentitalR = float64(squareAndMultiple(2,r, 10000000000))
      fmt.Println("exponentitalR is ", exponentitalR, " and R is ", r)
      } else {
        break
      }

    }

  r = r - 1
  exponentitalR = float64(squareAndMultiple(2,r, 10000000000))
  d := (varNumber/exponentitalR)



}
