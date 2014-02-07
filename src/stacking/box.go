package main

import (
	"fmt"
)

type box struct {
	H, L, W, Id int
}

var boxes []*box
var dp [][]int

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func stackHeight(below, usedBoxes int) int {
	if dp[below] == nil {
		dp[below] = make([]int, 1<<uint(len(boxes)/3))
	}
	//	fmt.Println(fmt.Sprintf("%v", boxes[below]))
	maxH := 0
	if dp[below][usedBoxes] == 0 {
		for idx, b := range boxes {
			if idx != below && !isBoxAlreadyUsed(usedBoxes, b) {
				if canPut(b, boxes[below]) {
//					fmt.Println(fmt.Sprintf("Yes %v %v %d", b, boxes[below], usedBoxes))
//					fmt.Println(fmt.Sprintf("%v", dp))
					sh := stackHeight(idx, usedBoxes|(1<<uint(b.Id)))
					maxH = max(maxH, sh+b.H)
					dp[below][usedBoxes] = maxH
					//				fmt.Println(maxH)
				}
			}
		}
	} else {
		maxH = dp[below][usedBoxes]
	}

	return maxH
}

func isBoxAlreadyUsed(usedBoxes int, b *box) bool {
	return (usedBoxes & (1 << uint(b.Id))) != 0
}

func canPut(above, below *box) bool {
	return (above.W <= below.W) && (above.L <= below.L)
}

func main() {
	var num int
	fmt.Scanf("%d", &num)
	boxes = make([]*box, 3*num)
	j := 0
	for i := 0; i < num; i++ {
		var l, w, h int
		fmt.Scanf("%d %d %d", &l, &w, &h)

		boxes[i+j] = &box{h, min(l, w), max(l, w), i}
		boxes[i+j+1] = &box{l, min(h, w), max(h, w), i}
		boxes[i+j+2] = &box{w, min(l, h), max(l, h), i}
		j += 2
	}

//	for _, b := range boxes {
//		fmt.Println(*b)
//	}
	var maxH int

	dp = make([][]int, 3*num)
	for idx, b := range boxes {
		maxH = max(maxH, stackHeight(idx, 1<<uint(b.Id))+b.H)
	}

	//		dp = make([][]int, 3*num)
	//		maxH = max(maxH, stackHeight(8, 1 << uint(boxes[8].Id)) + boxes[8].H)
	fmt.Println(maxH)
}
