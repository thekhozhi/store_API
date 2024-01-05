package controller

import (
	"fmt"
	"develop/models"
	"github.com/google/uuid"
)

func (c Controller) CreateProduct(){
product := getProductInfo()
id, err := c.Store.ProductStorage.Insert(product)
if err != nil{
	fmt.Println("Error while Inserting product!", err.Error())
}
fmt.Println("product's id is:", id)
}

func (c Controller) GetProductByID()  {
	idStr := ""
	fmt.Print("Enter id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil{
		fmt.Print("Id is not uuid error: ", err.Error())
		return
	}

	product, err := c.Store.ProductStorage.GetByID(id)
	if err != nil{
		fmt.Println("Error while getting product by id! :",err.Error())
		return
	}
	fmt.Println("the product is:", product)
}

func (c Controller) GetProductList(){
	products, err := c.Store.ProductStorage.GetList()
	if err != nil{
		fmt.Println("Error while getting list: ", err.Error())
		return
	}
	fmt.Println(products)
}

func (c Controller) UpdateProduct (){
	product := getProductInfo()

	err := c.Store.ProductStorage.Update(product)
	if err != nil{
		fmt.Println("Error while updating product: ", err.Error())
		return
	}
	if product.ID.String() != ""{
		fmt.Println("Successfully updated!")
	}else{
		fmt.Println("Successfullu created!")
	}
}



func getProductInfo() models.Product{
	var (
		idStr, name string
		cmd, price int
	)

	a:
		fmt.Print(`Enter command:
				1 - Create
				2 - Update
		`)
		fmt.Scan(&cmd)

		if cmd == 2 {
			fmt.Print("Enter id: ")
			fmt.Scan(&idStr)

			fmt.Println("Enter name: ")
			fmt.Scan(&name)

			fmt.Println("Enter price:")
			fmt.Scan(&price)
			
		}else if cmd == 1 {
			fmt.Println("Enter name: ")
			fmt.Scan(&name)

			fmt.Println("Enter price:")
			fmt.Scan(&price)

		}else{
			fmt.Println("Not found!")
			goto a
		}
		if idStr != ""{
			return models.Product{
				ID:		   uuid.MustParse(idStr),
				Name: 	   name,
				Price: 	   price,	
			}
		}
		return models.Product{
			 Name: name,
			 Price: price,
		}

}
