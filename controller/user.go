package controller

import (
	"fmt"
	"develop/models"
	"github.com/google/uuid"
)

func (c Controller) CreateUser(){
user := getUserInfo()
id, err := c.Store.UserStorage.Insert(user)
if err != nil{
	fmt.Println("Error while Inserting user!", err.Error())
}
fmt.Println("user's id is:", id)
}

func (c Controller) GetUserByID()  {
	idStr := ""
	fmt.Print("Enter id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil{
		fmt.Print("Id is not uuid error: ", err.Error())
		return
	}

	user, err := c.Store.UserStorage.GetByID(id)
	if err != nil{
		fmt.Println("Error while getting user by id! :",err.Error())
		return
	}
	fmt.Println("the user is:", user)
}

func (c Controller) GetUserList(){
	users, err := c.Store.UserStorage.GetList()
	if err != nil{
		fmt.Println("Error while getting list: ", err.Error())
		return
	}
	fmt.Println(users)
}

func (c Controller) UpdateUser (){
	user := getUserInfo()

	err := c.Store.UserStorage.Update(user)
	if err != nil{
		fmt.Println("Error while updating user: ", err.Error())
		return
	}
	if user.ID.String() != ""{
		fmt.Println("Successfully updated!")
	}else{
		fmt.Println("Successfullu created!")
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
				ID:		   uuid.MustParse(idStr),
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
