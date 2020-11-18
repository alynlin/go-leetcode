package main

import (
	"fmt"
	"sort"
)

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	ans := make([][]int, 0)
	if r0 > R || c0 > C {
		return nil
	}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			ans = append(ans, []int{i, j})
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		a, b := ans[i], ans[j]
		return abs(r0-a[0])+abs(c0-a[1]) < abs(r0-b[0])+abs(c0-b[1])
	})

	return ans
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func main() {
	ans := allCellsDistOrder(1, 2, 0, 0)

	fmt.Println(ans)

	ans = allCellsDistOrder(2, 2, 0, 1)
	fmt.Println(ans)
}

/**
输入：R = 2, C = 2, r0 = 0, c0 = 1
输出：[[0,1],[0,0],[1,1],[1,0]]
*/
