package postgres

import (
	"database/sql"
	"fmt"
	"develop/config"
	_ "github.com/lib/pq"
)

type Store struct {
	DB			  		*sql.DB
	UserStorage	  		userRepo
	OrderStorage        orderRepo
	ProductStorage 		productRepo
	OrderProductStorage orderProductRepo
}

func New(cfg config.Config) (Store, error){
	url := fmt.Sprintf(`host=%s port=%s user=%s password=%s database=%s`, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
	db, err := sql.Open("postgres", url)
	if err != nil{
		return Store{}, err
	}

	orderRepo := NewOrderRepo(db)
	productRepo := NewProductRepo(db)
	userRepo := NewUserRepo(db)
	orderProductRepo := NewOrderProductRepo(db)

	return Store{
		DB:		db,
		UserStorage: userRepo,
		OrderStorage: orderRepo,
		ProductStorage: productRepo,
		OrderProductStorage: orderProductRepo,
	}, nil
}