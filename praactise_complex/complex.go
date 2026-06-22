package main

import (
	"fmt"
)

type Complex struct{
	real int
	img int
}

func assignComplexNumber(real int,img int) *Complex{
	c:=Complex{
		real: real,
		img: img,
	}
	return &c
}

func (c *Complex) displayComplex(){
	fmt.Printf("%d + %di\n",c.real,c.img)
}

type ComplexInterface interface{
	additon(c Complex) Complex
	subtraction() Complex
}

func (c *Complex) additon(c1 Complex) *Complex{
	return  assignComplexNumber(c.real + c1.real,c.img + c1.img)
}

func (c *Complex) subtraction(c1 Complex) *Complex{
	return  assignComplexNumber(c.real - c1.real,c.img - c1.img)
}

func main(){
	new_complex1:=assignComplexNumber(3,4)
	new_complex2:=assignComplexNumber(2,5)

	fmt.Println("Complex Number 1")
	new_complex1.displayComplex()

	fmt.Println("Complex Number 2")
	new_complex2.displayComplex()
	
	sum:= new_complex1.additon(*new_complex2)
	diff:= new_complex1.subtraction(*new_complex2)

	fmt.Println("Addition: ")
	sum.displayComplex()

	fmt.Println("Subtraction: ")
	diff.displayComplex()
}