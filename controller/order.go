package controller

import (
	"fmt"
	"develop/models"
	"github.com/google/uuid"
)

func (c Controller) CreateOrder(){
order := getOrderInfo()
id, err := c.Store.OrderStorage.Insert(order)
if err != nil{
	fmt.Println("Error while Inserting order!", err.Error())
}
fmt.Println("order's id is:", id)
}

func (c Controller) GetOrderByID()  {
	idStr := ""
	fmt.Print("Enter id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil{
		fmt.Print("Id is not uuid error: ", err.Error())
		return
	}

	order, err := c.Store.OrderStorage.GetByID(id)
	if err != nil{
		fmt.Println("Error while getting order by id! :",err.Error())
		return
	}
	fmt.Println("the product is:", order)
}

func (c Controller) GetOrderList(){
	orders, err := c.Store.OrderStorage.GetList()
	if err != nil{
		fmt.Println("Error while getting list: ", err.Error())
		return
	}
	fmt.Println(orders)
}

func (c Controller) UpdateOrder (){
	order := getOrderInfo()

	err := c.Store.OrderStorage.Update(order)
	if err != nil{
		fmt.Println("Error while updating order: ", err.Error())
		return
	}
	if order.ID.String() != ""{
		fmt.Println("Successfully updated!")
	}else{
		fmt.Println("Successfullu created!")
	}
}



func getOrderInfo() models.Order{
	var (
		idStr, created_at string
		cmd, amount int
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

			fmt.Println("Enter amount: ")
			fmt.Scan(&amount)

			fmt.Println("Enter when was order created:")
			fmt.Scan(&created_at)
			
		}else if cmd == 1 {
			fmt.Println("Enter amount: ")
			fmt.Scan(&amount)

			fmt.Println("Enter created_at:")
			fmt.Scan(&created_at)

		}else{
			fmt.Println("Not found!")
			goto a
		}
		if idStr != ""{
			return models.Order{
				ID: uuid.MustParse(idStr),
				Amount: amount,
				CreatedAt: created_at,
			}
		}
		return models.Order{
			Amount: amount,
			CreatedAt: created_at,
		}
		 

}
