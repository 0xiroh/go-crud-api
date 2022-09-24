package odontologoUCInterface

import (
	odontologoModel "github.com/0xiroh/pkg/models/odontologo"
)

type OdontologoUCI interface {
	CreateData(string, string) error
	UpdateData(int64, string, string) error
	DeleteData(int64) error
	FindAll() ([]odontologoModel.Odontologos, error)
	FindByID(int64) (odontologoModel.Odontologos, error)
}