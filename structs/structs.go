// in go there is structs no classes
package main

import "fmt"

type Product struct{
	name string
	price int
	company string
}



func main(){


	
	p:=Product{
		name: "Samsung",
		price: 50000,
		company: "Samsung LLC",
	}
	fmt.Println("The product name is : ",p.name)
}




