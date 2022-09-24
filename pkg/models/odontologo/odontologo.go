package models

import (
	"github.com/jinzhu/gorm"
)

type Odontologos struct {
	gorm.Model
	Apellido string
	Nombre string
	Matricula string
}

type Odontologo struct {
	ID uint `json:"id"`
	Apellido string `json:"apellido"`
	Nombre string `json:"nombre"`
	Matricula string `json:"matricula"`
}