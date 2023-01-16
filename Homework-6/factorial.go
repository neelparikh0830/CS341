package main

import (
  "fmt"
  "math/big"
)

func fact(x *big.Int) *big.Int {
  yourInt := x.Uint64()
  if yourInt <= 1 {
    return big.NewInt(1)
  }
  return big.NewInt(1).Mul(x, fact(big.NewInt(1).Sub(x, big.NewInt(1))))
}

func main() {
  var input uint64
  fmt.Println("Enter the number: ")
  fmt.Scanf("%d", &input)
  value := new(big.Int).SetUint64(input)
  ans := fact(value)
  fmt.Println(ans)//fact(value))
}