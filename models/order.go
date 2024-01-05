package models

import "github.com/google/uuid"

type Order struct {
	ID 	  	   uuid.UUID
	Amount 	   int
	UserId    uuid.UUID
	CreatedAt string
}