// Copyright 2017 Venkatesh Gopal(vgopal3@jhu.edu), All rights reserved
package main

import (
  "fmt"
  "io/ioutil"
  "math/big"
  crypt "crypto/rand"
  "os"
)

func main() {

  if len(os.Args) != 3 {
    fmt.Println(" \n Follow command line specification \n ./rsa-keygen" +
      "<publickey-file-name> <privatekey-file-name>\n")

  } else {

  publickeyFileName := os.Args[1]
  privateKeyFileName := os.Args[2]

  p := getprimeNumber()
  q := getprimeNumber()

  N, publicKey, privateKey := rsaAlgorithmKeyGeneration(p,q)

  WritePublicKeyInformationToFile(N, publicKey, publickeyFileName)
  WritePrivateKeyInformationToFile(N, privateKey,p,q,privateKeyFileName )

  }
}

func WritePublicKeyInformationToFile(N *big.Int, publicKey *big.Int,
  publickeyFileName string) {

  NStringToWrite := N.String()
  leftBracket := "("
  rightBracket := ")"
  commaCharacter := ","
  publicKeyStringToWrite := publicKey.String()

  valueToWrite := leftBracket + NStringToWrite + commaCharacter +
  publicKeyStringToWrite + rightBracket

  err := ioutil.WriteFile(publickeyFileName, []byte(valueToWrite), 0644)
  if err != nil {
    fmt.Println("Some Problem in writing to a file")
  }

}

func WritePrivateKeyInformationToFile(N *big.Int, privateKey *big.Int, p *
  big.Int, q *big.Int, privateKeyFileName string) {

    NStringToWrite := N.String()
    commaCharacter := ","
    leftBracket := "("
    rightBracket := ")"
    privateKeyStringToWrite := privateKey.String()
    pStringToWrite := p.String()
    qStringToWrite := q.String()

    valueToWrite := leftBracket + NStringToWrite + commaCharacter +
     privateKeyStringToWrite + commaCharacter + pStringToWrite + commaCharacter+
     qStringToWrite + rightBracket

    err := ioutil.WriteFile(privateKeyFileName, []byte(valueToWrite), 0644)
    if err != nil {
      fmt.Println("Some Problem in writing to a file")
    }

}

func rsaAlgorithmKeyGeneration(p *big.Int, q *big.Int) (*big.Int,
  *big.Int, *big.Int) {

  // AS per RSA algorithm, the modulus is N = p.q

  N := big.NewInt(0)
  N = N.Mul(p,q)


  phiOfN := big.NewInt(0)
  pSub1 := (big.NewInt(0)).Sub(p,big.NewInt(1))
  qSub1 := (big.NewInt(0)).Sub(q,big.NewInt(1))
  phiOfN = phiOfN.Mul(pSub1,qSub1)

  e := generatePublicKey(phiOfN)


  eCopy := big.NewInt(0)
  eCopy = eCopy.Set(e)

  phiOfNCopy := big.NewInt(0)
  phiOfNCopy = phiOfNCopy.Set(phiOfN)
  // Testing Extended Euclidean Algorithm
  _,x,_ := extendedEuclideanAlgorithm(eCopy,phiOfNCopy)

  if (x.Cmp(big.NewInt(0)) == -1) {
    fmt.Println("Getting a negative value for X, so doing a mod operation")
    x = x.Add(x,phiOfN)

  }
  privateKey := big.NewInt(0)
  privateKey = privateKey.Set(x)


  temp := big.NewInt(0)
  temp = temp.Mul(e,privateKey)
  temp = temp.Mod(temp, phiOfN)

  return N, e, privateKey

}


func extendedEuclideanAlgorithm(a *big.Int, b *big.Int) (*big.Int,*big.Int,
*big.Int) {

  // Implementing the extendedEuclideanAlgorithm as per the pseudo-code
  // mentioned in the handbook of applied cryptography
  // http://cacr.uwaterloo.ca/hac/about/chap2.pdf (See Section 2.107)

  d := big.NewInt(0)
  x := big.NewInt(0)
  y := big.NewInt(0)

  if (b.Cmp(big.NewInt(0)) == 0) {

    d = d.Set(a)
    x = big.NewInt(1)
    y = big.NewInt(0)
    fmt.Println("First check return")
    return d,x,y
  }

  //  2 as per the Handbook of Applied cryptography
  x2 := big.NewInt(1)
  x1 := big.NewInt(0)
  y2 := big.NewInt(0)
  y1 := big.NewInt(1)

  // Setting big.Ints for the loop as we can't simple add (or) multiply
  // like Integers
  q := big.NewInt(0)
  r := big.NewInt(0)
  qb := big.NewInt(0)
  qx1 := big.NewInt(0)
  qy1 := big.NewInt(0)

  for ((b.Cmp(big.NewInt(0))) == 1) {

      // 3.1 as per the Handbook of Applied cryptography
      q = q.Div(a,b)
      r = r.Sub(a,qb.Mul(q,b))
      x = x.Sub(x2,qx1.Mul(q,x1))
      y = y.Sub(y2,qy1.Mul(q,y1))

      // 3.2 as per the Handbook of Applied cryptography

      a = a.Set(b)
      b = b.Set(r)
      x2 = x2.Set(x1)
      x1 = x1.Set(x)
      y2 = y2.Set(y1)
      y1 = y1.Set(y)
  }

  // 4 as per the Handbook of Applied cryptography

  d = d.Set(a)
  x = x.Set(x2)
  y = y.Set(y2)

  return d,x,y
}


func generatePublicKey(phiOfN *big.Int) (*big.Int) {


  e := big.NewInt(0)
   for true {
     e = getprimeNumber()

    phiOfNCopy := big.NewInt(0)
    phiOfNCopy = phiOfNCopy.Set(phiOfN)

    eCopy := big.NewInt(0)
    eCopy = eCopy.Set(e)

    gcd,_,_ := extendedEuclideanAlgorithm(eCopy,phiOfNCopy)

     if (gcd.Cmp(big.NewInt(1)) == 0) {
       break
     }
   }
   return e
}

func getprimeNumber()(*big.Int) {
    randomNumber := generateNumber()
  // Check for a prime number
  // I'm hardcoding the value of K in primality test to 5
    accuracyFactor := big.NewInt(5);
    resultWhetherPrime := false

    for (!resultWhetherPrime) {
        randomNumber = generateNumber()
        resultWhetherPrime = isaPrimeNumber(randomNumber,accuracyFactor)
        if (resultWhetherPrime) {
          break
        }
      }
      return randomNumber

}


func generateNumber() (*big.Int) {

  n := 64
  b := make([]byte, n)
  _, y := crypt.Read(b)
  if y != nil {
    fmt.Println("Some error")
  }

  z := big.NewInt(0)
  randomNumber := z.SetBytes(b)

  return randomNumber
}

func squareAndMultiple(a *big.Int, b *big.Int, c *big.Int) (*big.Int) {

  // FormatInt will provide the binary representation of a number
  binExp := fmt.Sprintf("%b", b)
  binExpLength := len(binExp)

  initialValue := big.NewInt(0)
  initialValue = initialValue.Mod(a,c)

  // Hold the initial value in result
  result := big.NewInt(0)
  result = result.Set(initialValue)

  // Using the square and multipy algorithm to perform modular exponentation
  for i := 1; i < binExpLength; i++ {

    // 49 is the ASCII representation of 1 and 48 is the ASCII representation
    // of 0
    interMediateResult := big.NewInt(0)
    interMediateResult = interMediateResult.Mul(result,result)
    result = result.Mod(interMediateResult, c)

    if byte(binExp[i]) == byte(49) {
      interResult := big.NewInt(0)
      interResult = interResult.Mul(result,initialValue)
      result = result.Mod(interResult, c)
    }
  }
  return result

}

func isaPrimeNumber(number *big.Int, accuracyFactor *big.Int) (bool) {

  // First finding the value of r, d as per equation ;
  // d * 2pow(r) = n - 1
  if (((big.NewInt(0)).Mod(number,big.NewInt(2))).Cmp(big.NewInt(0)) == 0) {
    // Case where the /dev/urandom has generated an even number
    return false
  } else {

  varNumber := (big.NewInt(0)).Sub(number, big.NewInt(1))

  r := big.NewInt(2)
  // exponentitalR is 2powr(r)
  exponentitalR := big.NewInt(2)

  for true {

    x := big.NewInt(0)
    modValForX := big.NewInt(0)
    x, modValForX = x.DivMod(varNumber, exponentitalR, modValForX)

    if ( (modValForX.Cmp(big.NewInt(0))) == 0) {
    // Fixing value 10000000000 for calculation purpose
    // To resue the squareAndMultiple algorithm but not affect the modulo part
      r = r.Add(r,big.NewInt(1))
      exponentitalR = squareAndMultiplyWithoutMod(big.NewInt(2),
      r)

      } else {
        break
      }

    }

  r = r.Sub(r,big.NewInt(1))

  exponentitalR = squareAndMultiplyWithoutMod(big.NewInt(2),
  r)

  d := big.NewInt(0)
  d = d.Div(varNumber,exponentitalR)

  for i := big.NewInt(0); (i.Cmp(accuracyFactor)) == -1;
  i.Add(i,big.NewInt(1)) {

  millerRabinPrimalityTestResult := millerRabinPrimalityTest(number, d,
  r)

  if (millerRabinPrimalityTestResult == false ) {
    return false
      }
    }
    return true
  }
}

func millerRabinPrimalityTest(number *big.Int, d *big.Int,
  r *big.Int) (bool) {

  // As per millerRabinPrimalityTest, we select an "a" in range[2,n-2]
  // Compute a value x = pow(a,d) % number and return true or false
  // based on some checks
  numberTemp := big.NewInt(0)
  numberTemp = (numberTemp.Sub(number, big.NewInt(4)))
  //aTemp := rand.Int63n(numberTemp.Int64()) + 2
  aTemp := int64(1000000000001)
  a := big.NewInt(aTemp)

  x := squareAndMultiple(a,d,number)

  numberMinusOne := (big.NewInt(0)).Sub(number, big.NewInt(1))
  if( ((x.Cmp(big.NewInt(1))) == 0) || ((x.Cmp(numberMinusOne)) == 0)) {
      return true
    }

  loopCount := (big.NewInt(0)).Sub(r,big.NewInt(1))

  for i := big.NewInt(0); (i.Cmp(loopCount)) == -1; i.Add(i,
    big.NewInt(1)) {

    xIntermediate := (big.NewInt(0)).Mul(x,x)

    x = x.Mod(xIntermediate,number)
    if (x.Cmp(big.NewInt(1)) == 0) {
      return false
    }
    if ((x.Cmp(numberMinusOne)) == 0) {
      return true
    }
  }
  return false

}

// Required since the previous function handles only exponentiation when I have
// a mod value

func squareAndMultiplyWithoutMod(number *big.Int, exponent *big.Int) (*big.Int){

	value := big.NewInt(1)
	// Below line to avail the binary
  // Operation performed later - > If 1, square and multiple
  // If 0, only square
	binExp := fmt.Sprintf("%b", exponent)
  binExpLength := len(binExp)

	if exponent == big.NewInt(1){
		return number
	}

	for i := 1; i < binExpLength; i++{
    // 49 is the ASCII representation of 1 and 48 is the ASCII representation
    // of 0
		if byte(binExp[i]) == byte(49){

      // temp := big.NewInt(0)
			value.Mul(value,value)
			value.Mul(value,number)

		}else{

      // temp := big.NewInt(0)
			value.Mul(value,value)

		}
	}

	return value

}
