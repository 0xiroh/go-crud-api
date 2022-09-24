package models

import (
	"github.com/jinzhu/gorm"
)

type Pacientes struct {
	gorm.Model
	Apellido string
	Nombre string
	Documento string
}

type Paciente struct {
	ID uint `json:"id"`
	Apellido string `json:"apellido"`
	Nombre string `json:"nombre"`
	Documento string `json:"matricula"`
}