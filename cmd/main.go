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
}