package main

import "fmt"
import "time"

func escrever(text string,canal chan string){
	for i:=0;i<5;i++{
		canal <- text
		time.Sleep(time.Second)
	}
}

func main(){
	canal := make(chan string)
	go escrever("ola mundo", canal)

	msg := <-canal
	fmt.Println(msg)
}