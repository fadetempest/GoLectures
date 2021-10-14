package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

func rNum() chan int{
	ch:=make(chan int)
	go func() {
		ch<-rand.Intn(21)
		close(ch)
	}()

	return ch
}

func calculate(ch ...<-chan int) <- chan int{
	cs:=make(chan int)
	wg:= sync.WaitGroup{}
	wg.Add(len(ch))
	for _,c:=range ch{
		go func(c <-chan int) {
			defer wg.Done()
			cs<-int(math.Pow(float64(<-c),2))
		}(c)
	}
	go func() {
		wg.Wait()
		close(cs)
	}()
	return cs
}

func main(){
	fanIn:=calculate(rNum(),rNum(),rNum(),rNum(),rNum(),rNum())
	for val:=range fanIn{
		fmt.Println(val)
	}
}
