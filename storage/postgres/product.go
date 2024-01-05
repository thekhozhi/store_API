package postgres

import (
	"database/sql"
	"fmt"
	"develop/models"

	"github.com/google/uuid"
)

type productRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) productRepo{
	return productRepo{
		DB: db,
	}
}
 

func (p productRepo) Insert(product models.Product)  (string, error) {
	id := uuid.New()

	_, err := p.DB.Exec(`INSERT INTO products values ($1, $2, $3)`, id, product.Name, product.Price)
	if err != nil{
		return "", err
	}
	return id.String(), nil
}

func (p productRepo) GetByID(id uuid.UUID) (models.Product, error){
product := models.Product{}

	err := p.DB.QueryRow(`SELECT from products where id = $1`, id).Scan(
		 &product.ID, &product.Name, &product.Price,
	)
	if err != nil{
		fmt.Println("Error while selecting product by id!", err.Error())
		return models.Product{}, err
	}
	return product, nil
}

func (p productRepo) GetList() ([]models.Product, error) {
	products := []models.Product{}

	rows, err := p.DB.Query(`SELECT * FROM products`)
	if err != nil{
		 return nil, err
	}
	 for rows.Next(){
		product := models.Product{}
		err := rows.Scan(&product.ID,  &product.Name, &product.Price)
		if err != nil{
			return nil, err
		}
		 products = append(products, product)
	 }
	 return products, nil
}

func (p productRepo) Update(product models.Product) error {
	_, err := p.DB.Exec(`UPDATE products set name = $1,  price = $2, where id = $3`,product.Name, product.Price, product.ID)
	if err != nil{
		return err
	}
	return nil
}