package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	//for {
	//	//	select {
	//	//	case ch <- 3:
	//	//		fmt.Println(<-ch)
	//	//		//close(ch)
	//	//	}
	//	//}
	//	//close(ch)

	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		fmt.Println("tがnil")
		fmt.Println("++++++++++")
		return
	}

	//fmt.Printf("t.Value1, %T: %v\n", t.Value, t.Value)
	fmt.Printf("t.Left, %T: %v\n", t.Left, t.Left)
	walk(t.Left, ch)
	fmt.Printf("t.Value, %T: %v\n", t.Value, t.Value)
	ch <- t.Value
	fmt.Printf("t.Right, %T: %v\n", t.Right, t.Right)
	walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
//func Same(t1, t2 *tree.Tree) bool

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
}