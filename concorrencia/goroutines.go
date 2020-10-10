package main

import "fmt"
import "time"

func escrever(text string){
	for{
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main(){
	go escrever("infinito") // goroutine
	escrever("dsds")
}