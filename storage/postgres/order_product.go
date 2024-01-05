package postgres

import (
	"database/sql"
	"fmt"
	"develop/models"

	"github.com/google/uuid"
)

type orderProductRepo struct {
	DB *sql.DB
}

func NewOrderProductRepo(db *sql.DB) orderProductRepo{
	return orderProductRepo{
		DB: db,
	}
}

func (op orderProductRepo) Insert(orderProduct models.OrderProduct)  (string, error) {
	id := uuid.New()

	_, err := op.DB.Exec(`INSERT INTO order_products values ($1, $2, $3)`, id, orderProduct.Quantity, orderProduct.Price)
	if err != nil{
		return "", err
	}
	return id.String(), nil
}

func (op orderProductRepo) GetByID(id uuid.UUID) (models.OrderProduct, error){
orderProduct := models.OrderProduct{}

	err := op.DB.QueryRow(`SELECT from order_products where id = $1`, id).Scan(
		 &orderProduct.ID, &orderProduct.OrderId, &orderProduct.ProductID, &orderProduct.Quantity, &orderProduct.Price)
	if err != nil{
		fmt.Println("Error while selecting order_products by id!", err.Error())
		return models.OrderProduct{}, err
	}
	return orderProduct, nil
}

func (op orderProductRepo) GetList() ([]models.OrderProduct, error) {
	orderProducts:= []models.OrderProduct{}

	rows, err := op.DB.Query(`SELECT * FROM order_products`)
	if err != nil{
		 return nil, err
	}
	 for rows.Next(){
		orderProduct := models.OrderProduct{}
		err := rows.Scan(&orderProduct.ID, &orderProduct.OrderId, &orderProduct.ProductID, &orderProduct.Quantity, &orderProduct.Price)
		if err != nil{
			return nil, err
		}
		 orderProducts = append(orderProducts, orderProduct)
	 }
	 return orderProducts, nil
}

func (op orderProductRepo) Update(orderProduct models.OrderProduct) error {
	_, err := op.DB.Exec(`UPDATE orders set quantity = $1,  price = $2 where id = $3`, orderProduct.Quantity, orderProduct.Price, orderProduct.ID)
	if err != nil{
		return err
	}
	return nil
}



