package main

import "fmt"

func main() {

	sl1 := []int{100, 911, 5, 0, 5, 1, 56, 2, -77, 8, 99, -2, 6, 988, 22, 33, 101, 8, -9, 7}
	quicksort(sl1, 0, len(sl1))
	fmt.Println(sl1)

}

var counter int = 0

func quicksort(sl1 []int, i int, le int) {
	pivot := sl1[le-1] // The pivot value is always the number at last index
	j := -1
	for ; i < le; i++ {
		if sl1[i] < pivot {
			if j > -1 {
				Swap(sl1, i, j)
				j++
			}
		} else {
			if j == -1 {
				j = i
			}
		}
		if i == le-1 {
			Swap(sl1, i, j)
		}
	}
	if j > 0 {
		quicksort(sl1, 0, j)
	}
	if j > 0 && j < le {
		quicksort(sl1, j+1, le)
	}

}

func Swap(sli []int, x int, y int) {
	sli[y], sli[x] = sli[x], sli[y]
}
