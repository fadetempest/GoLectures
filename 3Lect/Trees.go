package main

import (
	"golang.org/x/tour/tree"
	"fmt"
	)

var wg sync.WaitGroup
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	walking(t, ch)
	close(ch)
}

func walking(t *tree.Tree, ch chan int){
	if t != nil{
		walking(t.Left,ch)
		ch<-t.Value
		walking(t.Right,ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	ch1:=make(chan int)
	ch2:=make(chan int)
	
	go Walk(t1,ch1)
	go Walk(t2,ch2)
	
	value1:=<-ch1
	value2:=<-ch2
	
	if value1 != value2{
		return false
	}
	return true
}

func main() {
	ch:= make(chan int, 10)
	go Walk(tree.New(1), ch)
	for j:= range ch{
		fmt.Println(j)
	}
	fmt.Println(Same(tree.New(1),tree.New(1)))
	fmt.Println(Same(tree.New(1),tree.New(2)))
}
