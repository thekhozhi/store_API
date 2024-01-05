package postgres

import (
	"database/sql"
	"fmt"
	"develop/models"

	"github.com/google/uuid"
)

type userRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) userRepo{
	return userRepo{
		DB: db,
	}
}

func (u userRepo) Insert(user models.User)  (string, error) {
	id := uuid.New()

	_, err := u.DB.Exec(`INSERT INTO users values ($1, $2, $3, $4, $5)`, id, user.FirstName, user.LastName, user.Email, user.Phone)
	if err != nil{
		return "", err
	}
	return id.String(), nil
}

func (u userRepo) GetByID(id uuid.UUID) (models.User, error){
user := models.User{}

	err := u.DB.QueryRow(`SELECT from users where id = $1`, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
	)
	if err != nil{
		fmt.Println("Error while selecting user by id!", err.Error())
		return models.User{}, err
	}
	return user, nil
}

func (u userRepo) GetList() ([]models.User, error) {
	users := []models.User{}

	rows, err := u.DB.Query(`SELECT * FROM users`)
	if err != nil{
		 return nil, err
	}
	 for rows.Next(){
		user := models.User{}
		err := rows.Scan(&user.ID,  &user.FirstName, &user.LastName, &user.Email, &user.Phone)
		if err != nil{
			return nil, err
		}
		 users = append(users, user)
	 }
	 return users, nil
}

func (u userRepo) Update(user models.User) error {
	_, err := u.DB.Exec(`UPDATE users set first_name = $1, last_name = $2, email = $3, phone = $4 where id = $5`, 
	user.FirstName, user.LastName, user.Email, user.Phone, user.ID)
	if err != nil{
		return err
	}
	return nil
}