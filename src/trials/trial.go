package main 

import (
	"fmt"
	"sort"
)

//func sort(keys []string) []string{
//	if len(keys) <= 1{
//		return keys 
//	}
//	
//	pivotIdx := len(keys)/2
//	pivot := keys[pivotIdx]
//	var left, right []string
//	for idx, k := range keys{
//		if idx != pivotIdx{
//			if k < pivot{
//				left = append(left, k)
//			} else {
//				right = append(right, k)
//			}
//		}
//	}
//	
//	left = sort(left)
//	right = sort(right)
//	
//	var sorted []string
//	sorted = append(sorted, left...)
//	sorted = append(sorted, pivot)
//	sorted = append(sorted, right...)
//	return sorted
//	
//}

func min(a, b int) int{
	if a < b{ return a} else{ return b}
}

func mergeSort(keys []string) []string{
	size := len(keys)
	if size == 1{
		return keys
	}
	
	left := mergeSort(keys[0:size/2])
	right := mergeSort(keys[size/2:])
	
	return merge(left, right)
}

func merge(left, right []string) []string{
	leftLen := len(left)
	rightLen := len(right)
	
//	mergeSz := min(leftLen, rightLen)
	sorted := make([]string, leftLen + rightLen)
	
	for i,l,r := 0,0,0; i < (leftLen + rightLen); i++{
		if l >= leftLen{
			sorted[i] = right[r]
			r++
		} else if r >= rightLen{
			sorted[i] = left[l]
			l++
		} else if left[l] < right[r]{
			sorted[i] = left[l]
			l++
		} else{
			sorted[i] = right[r]
			r++
		}
	}
	
	return sorted
		
}


func main() {
	keys := []string{"aA", "Bb", "e", "c", "d", "ab", "a01", "a10", "kranthi", "chalasani"}
	sort.Strings(keys)
	fmt.Println(keys)
}

