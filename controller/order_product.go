package controller

import (
	"fmt"
	"develop/models"
	"github.com/google/uuid"
)

func (c Controller) CreateOrderProduct(){
orderProduct := getOrderProductInfo()
id, err := c.Store.OrderProductStorage.Insert(orderProduct)
if err != nil{
	fmt.Println("Error while Inserting order product!", err.Error())
}
fmt.Println("order product's id is:", id)
}

func (c Controller) GetOrderProductByID()  {
	idStr := ""
	fmt.Print("Enter id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil{
		fmt.Print("Id is not uuid error: ", err.Error())
		return
	}

	orderProduct, err := c.Store.OrderProductStorage.GetByID(id)
	if err != nil{
		fmt.Println("Error while getting order product by id! :",err.Error())
		return
	}
	fmt.Println("the product is:", orderProduct)
}

func (c Controller) GetOrderProductList(){
	orderProducts, err := c.Store.OrderProductStorage.GetList()
	if err != nil{
		fmt.Println("Error while getting list: ", err.Error())
		return
	}
	fmt.Println(orderProducts)
}

func (c Controller) UpdateOrderProduct (){
	orderProduct := getOrderProductInfo()

	err := c.Store.OrderProductStorage.Update(orderProduct)
	if err != nil{
		fmt.Println("Error while updating order product: ", err.Error())
		return
	}
	if orderProduct.ID.String() != ""{
		fmt.Println("Successfully updated!")
	}else{
		fmt.Println("Successfullu created!")
	}
}



func getOrderProductInfo() models.OrderProduct{
	var (
		idStr string
		cmd, quantity, price int
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

			fmt.Println("Enter quantity: ")
			fmt.Scan(&quantity)

			fmt.Println("Enter price:")
			fmt.Scan(&price)
			
		}else if cmd == 1 {
			fmt.Println("Enter quantity: ")
			fmt.Scan(&quantity)

			fmt.Println("Enter price:")
			fmt.Scan(&price)

		}else{
			fmt.Println("Not found!")
			goto a
		}
		if idStr != ""{
			return models.OrderProduct{
				 ID: uuid.MustParse(idStr),
				 Quantity: quantity,
				 Price: price,
			}
		}
		 return models.OrderProduct{
			Quantity:  quantity,
			Price:     price,
		}
}
