package main


func BubbleSort(input []int)[]int{
	array := input

	swapped := true;
	for swapped {
		swapped = false
		for i := 0; i < len(array) - 1; i++ {
			if array[i + 1] < array[i] {
				array[i], array[i+1] = array[i+1], array[i]
				swapped = true
			}
		}
	}

	return array
}


func CreateRandomArray(n int)[]int{
	result := []int{}
	for i:=0; i<n; i++{
		result = append(result, random(0,1000))
	}
	return result
}