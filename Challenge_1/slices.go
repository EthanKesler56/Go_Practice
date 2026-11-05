package main

func slice_test(test []int) []int {
	new_slice := make([]int, len(test))

	for i := 0; i < len(test); i++ {
		new_slice[i] = int(test[i])
		new_slice[i] = new_slice[i] * 2

	}
	return new_slice
}
