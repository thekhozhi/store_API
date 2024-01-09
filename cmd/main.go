package main

import (
	"develop/config"
	"develop/controller"
	"develop/storage/postgres"
	"fmt"
	"log"
	"net/http"
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
	http.HandleFunc("/user", con.User)
	http.HandleFunc("/product", con.Product)
	http.HandleFunc("/order", con.Order)
	http.HandleFunc("/order_product", con.OrderProduct)


	fmt.Println("listening at port :8080")
	http.ListenAndServe(":8080", nil)

	// USERS

	//con.CreateUser()
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