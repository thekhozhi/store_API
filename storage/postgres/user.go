package postgres

import (
	"database/sql"
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

func (u userRepo) GetByID(id string) (models.User, error){
user := models.User{}

	err := u.DB.QueryRow(`SELECT id, first_name, last_name, email, phone from users where id = $1`, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
	)
	if err != nil{
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

func (u userRepo) Delete(id string)error {
	id = "f3d55d0c-4213-41a7-a3fa-380b1f53e170"
	_, err := u.DB.Exec(`DELETE from users where id = $1`, id)
	if err != nil{
		return err
	}
	return nil
}