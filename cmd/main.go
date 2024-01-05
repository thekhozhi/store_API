package main

import (
	"log"
	"develop/config"
	"develop/controller"
	"develop/storage/postgres"
)

func main()  {
	cfg := config.Load()

	store, err := postgres.New(cfg)
	if err != nil{
		log.Fatalln("Error while connecting to db err: ", err.Error())
		return
	}
	defer store.DB.Close()

	 con := controller.New(store)

	// USERS

	 con.CreateUser()
	//  con.GetUserByID()
	//  con.GetUserList()
	//  con.UpdateUser()

	// PRODUCTS

	//  con.CreateProduct()
	//  con.GetProductByID()
	//  con.GetProductList()
	//  con.UpdateProduct()

	//ORDER PRODUCTS

	// con.CreateOrderProduct()
	// con.GetOrderProductByID()
	// con.GetOrderProductList()
	// con.UpdateOrderProduct()

	//ORDERS

	// con.CreateOrder()
	// con.GetOrderByID()
	// con.GetOrderList()
	// con.UpdateOrder()
}