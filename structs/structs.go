// in go there is structs no classes
package main

import "fmt"

type Product struct{
	name string
	price int
	company string
}

func (p *Product) display(){ //memebr function fo the product 
	fmt.Println("PRODUCT DETAILS :")
	fmt.Println("PRODUCT NAME : ",p.name)
	fmt.Println("PRODUCT COMPANY : ",p.company )
	fmt.Println("PRODUCT PRICE : ",p.price)
}

func newProduct(name string,price int,company string) *Product{
	p:= Product{
		name:name,
		price: price,
		company: company,
	}
	return &p 
}

func fun(copy_of_P Product){
	copy_of_P.name = "Iphone"
}

func main(){


	new_p := newProduct("Iphone",500,"Apple") // pointer to the prodct
	// fmt.Println("Product name : ",new_p.name)
	// fmt.Println("Product price : ",new_p.price)
	// fmt.Println("Product company : ",new_p.company)
	// p:=Product{
	// 	name: "Samsung",
	// 	price: 50000,
	// 	company: "Samsung LLC",
	// }
	// fmt.Println("The product name is : ",p.name)
// 	
}




