package postgres

import "develop/models"

type IStorage interface {
	Close()
	User() IUserStorage
	Product() IProductStorage
	Order() IOrderStorage
	OrderProduct() IOrderProductStorage
}

type IUserStorage interface {
	Create(models.User) (models.User, error)
	GetByID(id string) (models.User, error)
	GetList(models.User) ([]models.User, error)
	Update(models.User) (models.User, error)
	Delete(id string) error
}

type IProductStorage interface {
	Create(models.Product) (models.Product, error)
	GetByID(id string) (models.Product, error)
	GetList(models.Product) ([]models.Product, error)
	Update(models.Product) (models.Product, error)
	Delete(id string) error
}

type IOrderStorage interface {
	Create(models.Order) (models.Order, error)
	GetByID(id string) (models.Order, error)
	GetList(models.Order) ([]models.Order, error)
	Update(models.Order) (models.Order, error)
	Delete(id string) error
}

type IOrderProductStorage interface {
	Create(models.OrderProduct) (models.OrderProduct, error)
	GetByID(id string) (models.OrderProduct, error)
	GetList(models.OrderProduct) ([]models.OrderProduct, error)
	Update(models.OrderProduct) (models.OrderProduct, error)
	Delete(id string) error
}