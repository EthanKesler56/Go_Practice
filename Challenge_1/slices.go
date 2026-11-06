package main

import "fmt"

func slice_test(test []int) []int {
	new_slice := make([]int, len(test))

	for i := 0; i < len(test); i++ {
		new_slice[i] = int(test[i])
		new_slice[i] = new_slice[i] * 2

	}
	return new_slice
}

func slice_test_two() {
	names := []string{"Noel", "Nerissa", "Mori", "Shiori", "Kiara"}
	for i, elem := range names {
		fmt.Println(i, elem)
	}
}

func sliceOfSlices(rows, cols int) [][]int {
	sliceMatrix := [][]int{}
	for x := 0; x < rows; x++ {
		sliceMatrix = append(sliceMatrix, []int{})
		for j := 0; j < cols; j++ {
			sliceMatrix[x] = append(sliceMatrix[x], j+x)
		}
	}
	return sliceMatrix

}
