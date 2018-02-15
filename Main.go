package main

import (
	"math/rand"
	"fmt"

)



func main() {



	/////// BirthdayParty problem 1.4.44
	fmt.Println(DuplicateExpected(365))
	fmt.Println(RunTestDuplicate(365, 1000000))



	// Coupons problem 1.4.45
	fmt.Println(RunTestCoupons(365, 10000))




}







func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
