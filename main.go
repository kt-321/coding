package main

import (
	"fmt"
)

func Solution(A []int) int {
	var result int

	loop:
		for i:=1; i<1000001; i ++ {
			fmt.Printf("i: %v!!!\n", i)
			for n, v := range A {
				// 値が一致する時はループ抜ける
				if i == v {
					fmt.Printf("%vは値一致\n", v)
					break
				} else if n == len(A) - 1{
					// 最終ループで値が一致しない時は多重ループを停止
					result = i

					break loop
				}
				fmt.Println(i, v)
				fmt.Println("最終ループでなく値が一致しない")
			}
		}


	return result
}

func main() {
	s := []int {-1, -10, 1, 2, 100, 3, 4, 5}

	a := Solution(s)
	fmt.Println(a)
}
