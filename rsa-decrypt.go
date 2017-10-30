package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "math/big"
  "os"
)

func main() {

  if len(os.Args) != 3 {
    fmt.Println(" \n Follow command line specification \n ./rsa-decrypt" +
      "<privatekey-file-name> <Ciphertext to be decrypted, in decimal>\n")

  } else {

    file_name := os.Args[1]
    CipherTextInString := os.Args[2]

    N, privateKey, _, _ := ExtractDetailsFromPrivateKeyFile(file_name)

    Ciphertext := ConvertCipherTextToBigInt(CipherTextInString)
    Message := Decrypt(Ciphertext, N, privateKey)

    fmt.Println("Message is ", Message)
  }
}

func ConvertCipherTextToBigInt(CipherTextInString string) (*big.Int) {

  boolError := false
  Ciphertext := big.NewInt(0)

  Ciphertext, boolError = Ciphertext.SetString(CipherTextInString, 10)
  if boolError != true {
    fmt.Println(" Error in Set String")
    }

  return Ciphertext
}

func ExtractDetailsFromPrivateKeyFile(file_name string) (*big.Int, *big.Int,
  *big.Int, *big.Int) {

  FileContent, err := ioutil.ReadFile(file_name)
  N := big.NewInt(0)
  privateKey := big.NewInt(0)
  p := big.NewInt(0)
  q := big.NewInt(0)

  if err != nil {
    fmt.Println(" Error readng data from the file")
  } else {

  FileContentSliced := strings.Split(string(FileContent), ",")

  NinString := FileContentSliced[0]
  privateKeyComponentInString := FileContentSliced[1]
  pinString := FileContentSliced[2]
  qinString := FileContentSliced[3]

  boolError := false
  N, boolError = N.SetString(NinString,10)
  if boolError != true {
    fmt.Println(" Error in Set String")
    }

  privateKey, boolError = privateKey.SetString(privateKeyComponentInString, 10)
  if boolError != true {
    fmt.Println(" Error in Set String")
    }

  p, boolError = p.SetString(pinString, 10)
  if boolError != true {
    fmt.Println(" Error in Set String")
    }

  q, boolError = q.SetString(qinString, 10)
  if boolError != true {
    fmt.Println(" Error in Set String")
    }

  }

  return N, privateKey, p, q
}

func Decrypt(Ciphertext *big.Int, N *big.Int, privateKey *big.Int) (*big.Int){

  recoveredMessage := squareAndMultiple(Ciphertext, privateKey, N)
  return recoveredMessage

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
