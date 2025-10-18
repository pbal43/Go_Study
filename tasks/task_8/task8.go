package task_8

import (
	"errors"
	"fmt"
	"sort"
)

func Task8() {
	var numberSlice []int
	numberSlice = append(numberSlice, 1, 2, 5, 6, 2, 5, 8, 1, 9, 8, 3)
	sort.Ints(numberSlice)
	fmt.Println("Sorted slice: ", numberSlice)
	// Sorted slice:  [1 1 2 2 3 5 5 6 8 9]
	finder, err := binaryFinder(numberSlice, 0, len(numberSlice)-1, 9)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Index:", finder)
}

func binaryFinder(list []int, startIndex, finishIndex, number int) (int, error) {
	if finishIndex >= len(list) {
		return -1, errors.New("index out of range")
	}
	if len(list) == 0 {
		return -1, errors.New("list is empty")
	}
	if startIndex > finishIndex {
		return -1, errors.New("finish index must be bigger than start index (or NOT FOUND)")
	}
	middleIndex := startIndex + (finishIndex-startIndex)/2
	fmt.Println("Middle index: ", middleIndex)
	if list[middleIndex] == number {
		return middleIndex, nil
	}
	if list[middleIndex] > number {
		return binaryFinder(list, startIndex, middleIndex-1, number)
	}
	if list[middleIndex] < number {
		return binaryFinder(list, middleIndex+1, finishIndex, number)
	}
	return -1, errors.New("unexpected error")
}
