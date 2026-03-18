package main

import "fmt"

type Product struct{
	name string
	price int
}

func main(){
	products := make([]Product,0)

	products = append(products, Product{"Dell",1000})
	products = append(products, Product{"HP",5000})
	products = append(products, Product{"Dell Mouse",3000})
	products = append(products, Product{"RJ 45 Cable",500})
	fmt.Println(products)
}
