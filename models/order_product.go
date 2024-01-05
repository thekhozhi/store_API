package models

import "github.com/google/uuid"

type OrderProduct struct {
	ID 	      uuid.UUID
	OrderId   int
	ProductID uuid.UUID
	Quantity  int
	Price     int
 
}