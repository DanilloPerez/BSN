package main

import (
	"fmt"
	"strconv"
)

func main() {
	var BSN = ("243576891")
	VerifyBSN(BSN)
}

func generateBSN() {

}

func VerifyBSN(BSN string) bool {

	if len(BSN) != 9 {
		fmt.Println("YOU'RE ILLEGAL")
		return false
	}
	if ElevenTest(BSN) == false {
		fmt.Println("YOU'RE ILLEGAL")
		return false
	}

	fmt.Println("Seems valid to me, buddy")
	return true

}

func ElevenTest(BSN string) bool {
	var sum int = 0
	sum = sum - strconv.Atoi(BSN[8])
	for i := 0; i < 8; i++ {
		sum = sum + strconv.Atoi(BSN[i])*(9-i)
	}
	if sum%11 == 0 {
		return true
	}
	return false
}
