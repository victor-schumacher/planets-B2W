package entity

import "github.com/google/uuid"

type ID string

func NewID() ID {
	return ID(uuid.New().String())
}
