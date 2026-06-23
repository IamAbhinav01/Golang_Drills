package main

import (
	app "ComplexNumber/app"
	"fmt"
)

func main(){
	new_complex1:=app.AssignComplexNumber(3,4)
	new_complex2:=app.AssignComplexNumber(2,5)

	fmt.Println("Complex Number 1")
	new_complex1.DisplayComplex()

	fmt.Println("Complex Number 2")
	new_complex2.DisplayComplex()
	
	sum:= new_complex1.Addition(*new_complex2)
	diff:= new_complex1.Subtraction(*new_complex2)

	fmt.Println("Addition: ")
	sum.DisplayComplex()

	fmt.Println("Subtraction: ")
	diff.DisplayComplex()
}