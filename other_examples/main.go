package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Item struct {
	val   int
	subId int
}

type ItemHeap []Item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ItemHeap) Push(x any) {
	*h = append(*h, x.(Item))
}
func (h *ItemHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func divideAndSort(sli []int, n int) [][]int {
	var wg sync.WaitGroup
	var subs [][]int
	i := 0
	q := len(sli) / n
	r := len(sli) % n
	for i < len(sli) {
		wg.Add(1)
		l := q
		if r > 0 {
			r = r - 1
			l = l + 1
		}
		sub := sli[i : i+l]
		i = i + l
		go func() {
			defer wg.Done()
			fmt.Println(sub)
			sort.Slice(sub, func(i, j int) bool {
				return sub[i] < sub[j]
			})
			subs = append(subs, sub)
		}()
	}

	wg.Wait()
	return subs
}

func merge(subs [][]int, n int) []int {
	merged := make([]int, 0)
	pointers := make([]int, len(subs))
	h := make(ItemHeap, 0)
	heap.Init(&h)
	for i, sub := range subs {
		heap.Push(&h, Item{val: sub[0], subId: i})
		pointers[i] = 1
	}
	for h.Len() > 0 {
		it := heap.Pop(&h).(Item)
		merged = append(merged, it.val)
		p := pointers[it.subId]
		if p < len(subs[it.subId]) {
			heap.Push(&h, Item{val: subs[it.subId][p], subId: it.subId})
			pointers[it.subId] = p + 1
		}
	}
	return merged
}

func main() {

	for {
		sc := bufio.NewScanner(os.Stdin)
		fmt.Println("Please enter a sequence of integers, separated by space:")
		sc.Scan()
		arr := strings.Fields(sc.Text())
		sli := make([]int, 0, len(arr))
		for _, str := range arr {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Invalid input: %s\n", str)
				return
			}
			sli = append(sli, num)
		}

		subs := divideAndSort(sli, 4)
		merged := merge(subs, len(arr))

		fmt.Println()
		fmt.Println("The sorted sequence of integers are: ")
		fmt.Println(merged)
		fmt.Println()
	}

}
