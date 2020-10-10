package main

import "fmt"
import "time"
import "sync"

func escrever(text string){
	for i:=0;i<5;i++{
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main(){
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func(){
		escrever("infinito")
		waitGroup.Done()
	}()
	
	go func(){
		escrever("infinito 2")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}