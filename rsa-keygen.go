package main

import (
  "fmt"
  "math/rand"
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
  resultWhetherPrime := isaPrimeNumber(randomNumber,accuracyFactor)
  if (resultWhetherPrime) {
    fmt.Println("Got a prime number, in main")
  }
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

func isaPrimeNumber(number int, accuracyFactor int) (bool) {

  // First finding the value of r, d as per equation ;
  // d * 2pow(r) = n - 1
  if ((number % 2) == 0) {
    // Case where the /dev/urandom has generated an even number
    return false
  } else {
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

  for i := 0; i < accuracyFactor; i++ {
  millerRabinPrimalityTestResult := millerRabinPrimalityTest(number, int64(d),
  r)

  if (millerRabinPrimalityTestResult == false ) {
    return false
      }
    }
    return true
  }
}


func millerRabinPrimalityTest(number int, d int64, r int) (bool) {

  // As per millerRabinPrimalityTest, we select an "a" in range[2,n-2]
  // Compute a value x = pow(a,d) % number and return true or false
  // based on some checks
  a := rand.Int31n(int32(number) - 2)
  x := squareAndMultiple(int(a), int(d),number)

  if((x == 1) || (x == (number - 1))) {
      return true
    }

  loopCount := r -1
  for i := 0; i < loopCount; i++ {
    x = (x * x) % number
    if (x == 1) {
      return false
    }
    if (x == (number - 1)) {
      return true
    }
  }
  return false

}
