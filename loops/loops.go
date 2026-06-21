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

	 j:= 10
	for j> 0{
		if j == 7{
			fmt.Println("reached 3")
			j -- 
		}else if j == 5{
			fmt.Println("The number is 5")
			j--
			continue
		}else{

			fmt.Println("The value is descreasing",j)
			j--
		}
	}
} 
func main(){
	loops()
}