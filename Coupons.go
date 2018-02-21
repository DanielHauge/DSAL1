package main

import (
	"log"
)

func Couponons(n int)int{
	count := 0
	Remaining := n
	Chart := map[int]bool{}

	for {
		count++
		ChossenNumber := random(1, n+1)

		if Chart[ChossenNumber] == false{
			Chart[ChossenNumber] = true
			Remaining--
		}
		if Remaining == 0{
			break
			log.Println("No Remaning left")
		}
	}
	//log.Println(ThenFinalyCheck(Chart, n))
	return count
}

func ExpectedCoupons(n int)float64{


	H := 0.0

	for i := 1; i<=n; i++{
		H+=1/(float64(i))
	}

	return float64(n)*H;

}


func ThenFinalyCheck(Chart map[int]bool, n int)bool{

	for i := 1; i<n; i++{

		if Chart[i] == false{
			return false
		}

	}
	return true
}


func RunTestCoupons(n int, t int)float64{


	Results := []int{}

	for i:=t; i>0; i--{
		Results = append(Results, Couponons(n))
	}

	var total float64 = 0
	for _, value := range Results{
		total += float64(value)
	}
	return total/float64(len(Results))
}