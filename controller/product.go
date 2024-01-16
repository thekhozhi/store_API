package controller

import (
	"develop/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c Controller) Product(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateProduct(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetProductByID(w, r)
		} else {
			c.GetProductList(w, r)
		}
	case http.MethodPut:
		c.UpdateProduct(w, r)
	case http.MethodDelete:
		c.DeleteProduct(w, r)
	}
}

func (c Controller) CreateProduct(w http.ResponseWriter, r *http.Request){
	product := getProductInfo()
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		fmt.Println("error while reading data from client", err.Error())
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}
	
	product, err := c.Store.Product().Create(product)
	if err != nil{
		fmt.Println("Error while inserting product inside contrller err: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	
	
	resp, err := c.Store.Product().GetByID(product.ID)
	if err != nil{
		fmt.Println("Error while inserting product inside controller err: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
	}
	
	hanldeResponse(w, http.StatusOK, resp)
}
func (c Controller) GetProductByID(w http.ResponseWriter, r *http.Request)  {
	values := r.URL.Query()
	id := values["id"][0]


	product, err := c.Store.Product().GetByID(id)
	if err != nil{
		fmt.Println("Error while getting product by id! :",err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	hanldeResponse(w, http.StatusOK, product)
}

func (c Controller) GetProductList(w http.ResponseWriter, r *http.Request){
	products, err := c.Store.Product().GetList(models.Product{})
	if err != nil{
		fmt.Println("Error while getting list: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	hanldeResponse(w, http.StatusOK, products)
}

func (c Controller) UpdateProduct (w http.ResponseWriter, r *http.Request){
	product := getProductInfo()

	product,err := c.Store.Product().Update(product)
	if err != nil{
		fmt.Println("Error while updating product: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	if product.ID != ""{
		fmt.Println("Successfully updated!")
		hanldeResponse(w, http.StatusOK, product)
	}else{
		fmt.Println("Successfullu created!")
		hanldeResponse(w, http.StatusOK, product)
	}
}

func (c Controller) DeleteProduct(w http.ResponseWriter, r *http.Request){
	id := "bd5e0a1c-c37b-405a-8ec8-3430746c86a3"
	err := c.Store.Product().Delete(id)
	if err != nil{
		fmt.Println("Error while deleting user!")
		return
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
				ID:		   idStr,
				Name: 	   name,
				Price: 	   price,	
			}
		}
		return models.Product{
			 Name: name,
			 Price: price,
		}

}
