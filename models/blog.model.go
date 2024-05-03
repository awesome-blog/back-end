package models

import "github.com/google/uuid"

type BlogModel struct {
	Id   uuid.UUID `json:"id,omitempty" bson:"_id,omitempty"`
	Path string    `json:"path,omitempty"`
}
