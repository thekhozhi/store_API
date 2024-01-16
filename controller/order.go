package controller

import (
	"develop/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c Controller) Order(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateOrder(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetOrderByID(w, r)
		} else {
			c.GetOrderList(w, r)
		}
	case http.MethodPut:
		c.UpdateOrder(w, r)
	case http.MethodDelete:
		c.DeleteOrder(w,r)
	}
}

func (c Controller) CreateOrder(w http.ResponseWriter, r *http.Request){
	order := getOrderInfo()
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		fmt.Println("error while reading data from client", err.Error())
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}

	order, err := c.Store.Order().Create(order)
if err != nil{
	fmt.Println("Error while inserting order inside controller err: ", err.Error())
	hanldeResponse(w, http.StatusInternalServerError, err)
	return
}


resp, err := c.Store.Order().GetByID(order.ID)
if err != nil{
	fmt.Println("Error while inserting order inside controller err: ", err.Error())
	hanldeResponse(w, http.StatusInternalServerError, err)
}

hanldeResponse(w, http.StatusOK, resp)
}

func (c Controller) GetOrderByID(w http.ResponseWriter, r *http.Request)  {
	values := r.URL.Query()
	id := values["id"][0]


	order, err := c.Store.Order().GetByID(id)
	if err != nil{
		fmt.Println("Error while getting order by id! :",err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	hanldeResponse(w, http.StatusOK, order)
}

func (c Controller) GetOrderList(w http.ResponseWriter, r *http.Request){
	orders, err := c.Store.Order().GetList(models.Order{})
	if err != nil{
		fmt.Println("Error while getting list: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	hanldeResponse(w, http.StatusOK, orders)
}



func (c Controller) UpdateOrder (w http.ResponseWriter, r *http.Request){
	order := getOrderInfo()

	order,err := c.Store.Order().Update(order)
	if err != nil{
		fmt.Println("Error while updating orders: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	if order.ID != ""{
		fmt.Println("Successfully updated!")
		hanldeResponse(w, http.StatusOK, order)
	}else{
		fmt.Println("Successfullu created!")
		hanldeResponse(w, http.StatusOK, order)
	}
}

func (c Controller) DeleteOrder(w http.ResponseWriter, r *http.Request){
	id := "1"
	err := c.Store.Product().Delete(id)
	if err != nil{
		fmt.Println("Error while deleting order!")
		return
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
				ID: idStr,
				Amount: amount,
				CreatedAt: created_at,
			}
		}
		return models.Order{
			Amount: amount,
			CreatedAt: created_at,
		}
		 

}
