package utils

import (

)


const (
	MaxUint = ^uint(0) 
	MinUint = 0 
	MaxInt = int(MaxUint >> 1) 
	MinInt = -MaxInt - 1
)


func MinInts(x, y int) int {
    if x < y {
        return x
    }
    return y
}
func MaxInts(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func MinIntsSlice(s []int) int {
	N := len(s)
	min := MaxInt

	for i := 0; i < N; i++ {
		if (s[i] < min) { min = s[i] }
	}

	return min
}
func MaxIntsSlice(s []int) int {
	N := len(s)
	max := MinInt
	
	for i := 0; i < N; i++ {
		if (s[i] > max) { max = s[i] }
	}

	return max
}

func MakeSliceInt(min, max, increment int) []int {
	N := (max - min) / increment + 1
	result := make([]int, N)
	for i := 0; i < N; i++ {
		result[i] = min + i * increment
	}

	return result
}

