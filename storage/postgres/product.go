package postgres

import (
	"database/sql"
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

func (p productRepo) GetByID(id string) (models.Product, error){
product := models.Product{}

	err := p.DB.QueryRow(`SELECT id, name, price from products where id = $1`, id).Scan(
		 &product.ID, &product.Name, &product.Price,
	)
	if err != nil{
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

func (p productRepo) Delete(id string)error {
	id =  "bd5e0a1c-c37b-405a-8ec8-3430746c86a3"
	_, err := p.DB.Exec(`DELETE from products where id = $1`, id)
	if err != nil{
		return err
	}
	return nil
}