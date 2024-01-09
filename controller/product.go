package controller

import (
	"develop/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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
		// delete
	}
}

func (c Controller) CreateProduct(w http.ResponseWriter, r *http.Request){
	product := getProductInfo()
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		fmt.Println("error while reading data from client", err.Error())
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}
	
	id, err := c.Store.ProductStorage.Insert(product)
	if err != nil{
		fmt.Println("Error while inserting product inside contrller err: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	
	
	resp, err := c.Store.ProductStorage.GetByID(uuid.MustParse(id))
	if err != nil{
		fmt.Println("Error while inserting product inside controller err: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
	}
	
	hanldeResponse(w, http.StatusOK, resp)
}
func (c Controller) GetProductByID(w http.ResponseWriter, r *http.Request)  {
	values := r.URL.Query()
	id := values["id"][0]


	product, err := c.Store.ProductStorage.GetByID(uuid.MustParse(id))
	if err != nil{
		fmt.Println("Error while getting product by id! :",err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	hanldeResponse(w, http.StatusOK, product)
}

func (c Controller) GetProductList(w http.ResponseWriter, r *http.Request){
	products, err := c.Store.ProductStorage.GetList()
	if err != nil{
		fmt.Println("Error while getting list: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	hanldeResponse(w, http.StatusOK, products)
}

func (c Controller) UpdateProduct (w http.ResponseWriter, r *http.Request){
	product := getProductInfo()

	err := c.Store.ProductStorage.Update(product)
	if err != nil{
		fmt.Println("Error while updating product: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	if product.ID.String() != ""{
		fmt.Println("Successfully updated!")
		hanldeResponse(w, http.StatusOK, product)
	}else{
		fmt.Println("Successfullu created!")
		hanldeResponse(w, http.StatusOK, product)
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
