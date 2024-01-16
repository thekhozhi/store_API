package controller

import (
	"develop/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c Controller) User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateUser(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetUserByID(w, r)
		} else {
			c.GetUserList(w, r)
		}
	case http.MethodPut:
		c.UpdateUser(w, r)
	case http.MethodDelete:
		c.DeleteUser(w, r)
	}
}

func (c Controller) CreateUser(w http.ResponseWriter, r *http.Request){
user := getUserInfo()
if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
	fmt.Println("error while reading data from client", err.Error())
	hanldeResponse(w, http.StatusBadRequest, err)
	return
}

user, err := c.Store.User().Create(user)
if err != nil{
	fmt.Println("Error while inserting user inside contrller err: ", err.Error())
	hanldeResponse(w, http.StatusInternalServerError, err)
	return
}


resp, err := c.Store.User().GetByID(user.ID)
if err != nil{
	fmt.Println("Error while inserting user inside controller err: ", err.Error())
	hanldeResponse(w, http.StatusInternalServerError, err)
}

hanldeResponse(w, http.StatusOK, resp)
}

func (c Controller) GetUserByID(w http.ResponseWriter, r *http.Request)  {
	values := r.URL.Query()
	id := values["id"][0]


	user, err := c.Store.User().GetByID(id)
	if err != nil{
		fmt.Println("Error while getting user by id! :",err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	hanldeResponse(w, http.StatusOK, user)
}

func (c Controller) GetUserList(w http.ResponseWriter, r *http.Request){
	users, err := c.Store.User().GetList(models.User{})
	if err != nil{
		fmt.Println("Error while getting list: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	hanldeResponse(w, http.StatusOK, users)
}

func (c Controller) UpdateUser (w http.ResponseWriter, r *http.Request){
	user := getUserInfo()

	user,err := c.Store.User().Update(user)
	if err != nil{
		fmt.Println("Error while updating user: ", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}
	if user.ID != ""{
		fmt.Println("Successfully updated!")
		hanldeResponse(w, http.StatusOK, user)
	}else{
		fmt.Println("Successfullu created!")
		hanldeResponse(w, http.StatusOK, user)
	}
}

func (c Controller) DeleteUser(w http.ResponseWriter, r *http.Request){
	id := "fbd91626-7594-4d33-b4be-6d6bf0c15a98"
	err := c.Store.User().Delete(id)
	if err != nil{
		fmt.Println("Error while deleting user!")
		return
	}
}



func getUserInfo() models.User{
	var (
		idStr, firstName, lastName, email, phone string
		cmd int
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

			fmt.Println("Enter firstname and lastname: ")
			fmt.Scan(&firstName, &lastName)

			fmt.Println("Enter email:")
			fmt.Scan(&email)

			fmt.Println("Enter phone number:")
			fmt.Scan(&phone)
			
		}else if cmd == 1 {
			fmt.Println("Enter firstname and lastname: ")
			fmt.Scan(&firstName, &lastName)

			fmt.Println("Enter email:")
			fmt.Scan(&email)

			fmt.Println("Enter phone number:")
			fmt.Scan(&phone)

		}else{
			fmt.Println("Not found!")
			goto a
		}
		if idStr != ""{
			return models.User{
				ID:		   idStr,
				FirstName: firstName,
				LastName:  lastName,
				Email: 	   email,
				Phone:     phone,
			}
		}
		return models.User{
			FirstName: firstName,
			LastName:  lastName,
			Email: 	   email,
			Phone:     phone,
		}

}
