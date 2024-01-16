package postgres

import (
	"database/sql"
	"develop/models"

	"github.com/google/uuid"
)

type orderRepo struct {
	DB *sql.DB
}

func NewOrderRepo(db *sql.DB) orderRepo{
	return orderRepo{
		DB: db,
	}
}

func (o orderRepo) Insert(order models.Order)  (string, error) {
	id := uuid.New()

	_, err := o.DB.Exec(`INSERT INTO orders values ($1, $2, $3, $4)`, id, order.Amount, order.UserId, order.CreatedAt)
	if err != nil{
		return "", err
	}
	return id.String(), nil
}

func (o orderRepo) GetByID(id string) (models.Order, error){
order := models.Order{}

	err := o.DB.QueryRow(`SELECT from orders where id = $1`, id).Scan(
		 &order.ID, &order.Amount, &order.UserId, &order.CreatedAt,)
	if err != nil{
		return models.Order{}, err
	}
	return order, nil
}

func (o orderRepo) GetList() ([]models.Order, error) {
	orders:= []models.Order{}

	rows, err := o.DB.Query(`SELECT * FROM orders`)
	if err != nil{
		 return nil, err
	}
	 for rows.Next(){
		order := models.Order{}
		err := rows.Scan(&order.ID,  &order.Amount, & order.UserId, &order.CreatedAt)
		if err != nil{
			return nil, err
		}
		 orders = append(orders, order)
	 }
	 return orders, nil
}

func (o orderRepo) Update(order models.Order) error {
	_, err := o.DB.Exec(`UPDATE orders set amount = $1,  user_id = $2, set = created_at = $3 where id = $4`, order.Amount, order.UserId, order.CreatedAt, order.ID)
	if err != nil{
		return err
	}
	return nil
}

func (o orderRepo) Delete(id string)error {
	id =  "1"
	_, err := o.DB.Exec(`DELETE from orders where id = $1`, id)
	if err != nil{
		return err
	}
	return nil
}


