package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Lets Play with Math")

	// var myNum1 int = 2
	// var myNum2 float64 = 4

	// fmt.Println("The sum is: ", myNum1+int(myNum2))

	//random number

	//Through math
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	//Through Crypto
	myrandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myrandomNumber)

}
