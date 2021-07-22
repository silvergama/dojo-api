package entity

import "github.com/google/uuid"

// ID entity ID
type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}
