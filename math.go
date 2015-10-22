package utils

import (

)


const (
	MaxUintConst = ^uint(0) 
	MinUintConst = 0 
	MaxIntConst = int(MaxUint >> 1) 
	MinIntConst = -MaxInt - 1
)


func MinInt(x, y int) int {
    if x < y {
        return x
    }
    return y
}
func MaxInt(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func MinIntSlice(s []int) int {
	N := len(s)
	min := MaxIntConst

	for i := 0; i < N; i++ {
		if (s[i] < min) { min = s[i] }
	}

	return min
}
func MaxIntSlice(s []int) int {
	N := len(s)
	max := MinIntConst
	
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

