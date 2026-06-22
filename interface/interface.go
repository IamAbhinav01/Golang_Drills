package main

import "fmt"

type Product struct{
	name string
	price int
	company string
}

type SellableProduct interface{
	buy() 
	getDiscount() int
}

func  (p *Product)buy(){
	fmt.Println("Buying the Product : ",p.name," for price : ",p.price," by the company : ",p.company)

}
func  (p *Product)getDiscount() int {
	discount:=p.price * 20/100
	fmt.Println("Discount of the product is : ",discount)
	return discount

}

func newProduct(name string,price int,company string) *Product{
	p:= Product{
		name:name,
		price: price,
		company: company,
	}
	return &p 
}

func checkDiscount(product SellableProduct){
	discount:=product.getDiscount()
	if discount > 50{
		fmt.Println("Discoutn is good , buying the product")
		product.buy()
		return
	}else{
		fmt.Print("Discount is not good , not buying the product")
		return 
	}
}


func main(){
	new_p := newProduct("Iphone",500,"Apple")
	checkDiscount(new_p)

}