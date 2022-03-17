package main

import (
	"fmt"
)

/*

 */
var testCases = []struct {
	nums   []int
	k      int
	result int
}{
	{[]int{2, 1}, 2, 1},
}

func main() {
	for _, tc := range testCases {
		got := 0
		want := tc.result
		if got != want {
			fmt.Printf("got: %v, want: %v\n", got, want)
		} else {
			fmt.Println("ok")
		}
	}

	//for _, tc := range testCases {
	//	got := searchRange(tc.nums, tc.target)
	//	want := tc.result
	//	ok := true
	//	for i := range want {
	//		if got[i] != want[i] {
	//			ok = false
	//			break
	//		}
	//	}
	//
	//	if ok {
	//		fmt.Println("ok")
	//	} else {
	//		fmt.Printf("got: %d, want: %d\n", got, want)
	//	}
	//}
}
