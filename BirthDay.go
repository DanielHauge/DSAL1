package main

import "math"

func RunTestDuplicate(n int, t int)float64{


	Results := []int{}

	for i:=t; i>0; i--{
		Results = append(Results, NumberDuplicate(n))
	}

	var total float64 = 0
	for _, value := range Results{
		total += float64(value)
	}
	return total/float64(len(Results))
}


func DuplicateExpected(n int)float64{
	return math.Sqrt(math.Pi*float64(n)/2)
}

func NumberDuplicate(n int)int{
	count := 0;
	Chart := map[int]bool{}

	for i := n-1; i>0; i--{
		count++
		ChossenNumber := random(1, n+1)
		if Chart[ChossenNumber] == true{
			break
		}else {
			Chart[ChossenNumber] = true
		}
	}
	return count
}
