package controller

import (
	"develop/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (c Controller) OrderProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateOrderProduct(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetOrderProductByID(w, r)
		} else {
			c.GetOrderProductList(w, r)
		}
	case http.MethodPut:
		c.UpdateOrderProduct(w, r)
	case http.MethodDelete:
		// delete
	}
}

func (c Controller) CreateOrderProduct(w http.ResponseWriter, r *http.Request){
	orderProduct := getOrderProductInfo()
	if err := json.NewDecoder(r.Body).Decode(&orderProduct); err != nil {
		fmt.Println("error while reading data from client", err.Error())
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}
	
	id, err := c.Store.OrderProductStorage.Insert(orderProduct)
	if err != nil{
		fmt.Println("Error while inserting order product inside contrller err: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	
	
	resp, err := c.Store.OrderProductStorage.GetByID(uuid.MustParse(id))
	if err != nil{
		fmt.Println("Error while inserting order product inside controller err: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
	}
	
	hanldeResponse(w, http.StatusOK, resp)
	}
	

func (c Controller) GetOrderProductByID(w http.ResponseWriter, r *http.Request)  {
	values := r.URL.Query()
	id := values["id"][0]


	orderProduct, err := c.Store.OrderProductStorage.GetByID(uuid.MustParse(id))
	if err != nil{
		fmt.Println("Error while getting order product by id! :",err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	hanldeResponse(w, http.StatusOK, orderProduct)
}

func (c Controller) GetOrderProductList(w http.ResponseWriter, r *http.Request){
	orderProducts, err := c.Store.OrderProductStorage.GetList()
	if err != nil{
		fmt.Println("Error while getting list: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	hanldeResponse(w, http.StatusOK, orderProducts)
}


func (c Controller) UpdateOrderProduct (w http.ResponseWriter, r *http.Request){
	orderProduct := getOrderProductInfo()

	err := c.Store.OrderProductStorage.Update(orderProduct)
	if err != nil{
		fmt.Println("Error while updating order products: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	if orderProduct.ID.String() != ""{
		fmt.Println("Successfully updated!")
		hanldeResponse(w, http.StatusOK, orderProduct)
	}else{
		fmt.Println("Successfullu created!")
		hanldeResponse(w, http.StatusOK, orderProduct)
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
