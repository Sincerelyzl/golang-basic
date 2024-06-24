package main

import "fmt"

func main() {
	// var a [3]int = [3]int{1, 2, 3} //arrays
	// fmt.Println(a[0], a[1], a[2])

	// s := []int{1, 2, 3} //slice
	// fmt.Println(s)

	arr1 := [5]int{1, 2, 3, 4, 5}
	// arr2 := [4]int{4, 5, 6, 20}
	// fmt.Println(sumArray(arr))
	// fmt.Println(concatArrays(arr1, arr2))

	s := []int{1, 2, 3, 4, 5}
	// fmt.Println(averageSlice(s))
	// fmt.Println(maxInSlice(s))
	// fmt.Println(filterEvenNumbers(s))

	// fmt.Println(mergeArrayAndSlice(arr1, s))

	fmt.Println(incrementArray(arr1))

	if contain(s, 6) {
		fmt.Println("Have this number")
	} else {
		fmt.Println("Not have this number")
	}

}

func sumArray(arr [5]int) int {

	sum := 0

	for i := range arr {
		sum += arr[i]
	}
	return sum

}

func averageSlice(s []int) float64 {

	sum := 0

	for i := range s {
		sum += s[i]
	}
	avg := sum / len(s)

	return float64(avg)
}

func maxInSlice(s []int) int {
	var max int

	for i := range s {
		if max == 0 {
			max = s[i]
			continue
		}
		if s[i] > max {
			max = s[i]
		}
	}

	return max
}

func filterEvenNumbers(s []int) []int {
	evenSlice := []int{}

	for i := range s {
		if s[i]%2 == 0 {
			evenSlice = append(evenSlice, s[i])
		}
	}
	return evenSlice

}

func concatArrays(arr1 [4]int, arr2 [4]int) []int {
	fusionSlice := []int{}

	for i := range arr1 {
		fusionSlice = append(fusionSlice, arr1[i])
	}
	for i := range arr2 {
		fusionSlice = append(fusionSlice, arr2[i])
	}
	return fusionSlice
}

func mergeArrayAndSlice(arr [4]int, s []int) []int {
	mergedArraySlice := []int{}

	for i := range arr {
		mergedArraySlice = append(mergedArraySlice, arr[i])
	}
	for i := range s {
		mergedArraySlice = append(mergedArraySlice, s[i])
	}
	return mergedArraySlice

}

func incrementArray(arr [5]int) [5]int {
	plusUpArr := [5]int{}
	for i := range arr {
		plusUpArr[i] = arr[i] + 1
	}
	return plusUpArr
}

func contain(s []int, value int) bool {
	checker := false

	for i := range s {
		if s[i] == value {
			checker = true
		}
	}
	return checker
}
