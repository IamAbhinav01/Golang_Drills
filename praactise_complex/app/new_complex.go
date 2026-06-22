package app

import (
	"fmt"
)

type Complex struct{
	Real int
	Img int
}

func AssignComplexNumber(real int,img int) *Complex{
	c:=Complex{
		Real: real,
		Img: img,
	}
	return &c
}

func (c *Complex) DisplayComplex(){
	fmt.Printf("%d + %di\n",c.Real,c.Img)
}

func (c *Complex) Addition(c1 Complex) *Complex{
	return AssignComplexNumber(c.Real + c1.Real,c.Img + c1.Img)
}

func (c *Complex) Subtraction(c1 Complex) *Complex{
	return AssignComplexNumber(c.Real - c1.Real,c.Img - c1.Img)
}
