package app

import (
	"github.com/Firgisotya/go-rest-api/app/models"
)

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Category{}},
		{Model: models.Books{}},
	}
}