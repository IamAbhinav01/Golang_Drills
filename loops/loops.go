package main

import "fmt"

func loops(){
	for  i := 0;i<5;i++{
		fmt.Println("loop",i)
	}

	for i:= range 10{
		fmt.Println("Range is ",i)
	}
	for i,char:= range "Abhinav"{
		fmt.Println("Range of Abhinav",i,char)
	}
} 
func main(){
	loops()
}