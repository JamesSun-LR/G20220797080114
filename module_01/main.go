package main

import (
	"fmt"
	"time"
)

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
		fmt.Printf("Put data: %d\n", i)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for k := range ch {
		time.Sleep(1 * time.Second)
		fmt.Printf("Get data: %d\n", k)
	}
}

func main() {
	fmt.Println("==== Task 1.1 ====")
	str := [5]string{"I", "am", "stupid", "adn", "weak"}
	for i := range str {
		if str[i] == "stupid" {
			str[i] = "smart"
		} else if str[i] == "weak" {
			str[i] = "strong"
		}
	}
	fmt.Println(str)

	fmt.Println("==== Task 1.2 ====")
	ch := make(chan int, 3)
	go produce(ch)
	time.Sleep(5 * time.Second)
	consumer(ch)
}
