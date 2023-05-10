package algorithms

import "math/rand"

func GenerateRandomSlice(n int) []int {
	arr := rand.Perm(n)
	for i, val := range arr {
		arr[i] = val + 1
	}
	return arr
}
