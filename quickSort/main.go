package main

import (
	"fmt"
)

func main() {
	arr := []int{8, 7, 5, 3, 2, 3, 5, 21}
	qSort(arr, 0, len(arr) - 1)
	fmt.Println(arr)
}

func qSort(arr []int, from int, to int) {
	if from >= to {
		return
	}
	flag := arr[from]
	left := from
	right := to
	for left < right {
		for arr[right] >= flag && left < right {
			right--
		}
		for arr[left] <= flag && left < right {
			left++
		}
		if (left < right) {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[from], arr[left] = arr[left], arr[from]
	qSort(arr, from, left-1)
	qSort(arr, left+1, to)
}
