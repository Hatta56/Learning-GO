package main

import "time"

// func sum(s []int, c chan int){
// 	sum := 0
// 	for _, v range s {
// 		sum += v
// 	}
// 	c <- sum
// }

// func main(){
// 	s:= []int{7,2,8,-9,4,0}

// 	c:= make(chan int)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)
// 	x, y := <-c, <-c
// 	fmt.Println(x, y, x+y)
// }

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		println(s)
	}
}

func main() {
	go say("world")
}
