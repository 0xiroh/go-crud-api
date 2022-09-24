package repository

import (
	odontologoModel "github.com/0xiroh/go-crud-api/pkg/models/odontologo"
	odontologoI "github.com/0xiroh/go-crud-api/pkg/interface/odontologoInterface"
	"errors"

	"github.com/jinzhu/gorm"
)
import (
	
)
type dbConn struct {
	DbConn *gorm.DB
}

// NewOdontologoRepo - initial repo block name
func NewOdontologoRepo(DbConn *gorm.DB) odontologoI.OdontologoI {
	return &dbConn{DbConn}
}

// CreateData - create data in table odontologo_s
func (config *dbConn) CreateData(nombre, apellido, matricula string) error {
	myData := odontologoModel.Odontologo{
		Nombre:        nombre,
		Apellido: apellido,
		Matricula: matricula,
	}

	err := config.DbConn.Create(&myData).Error
	if err != nil {
		return errors.New("CreateData err = " + err.Error())
	}

	return nil
}

// UpdateData - update data in table odontologo_s
func (config *dbConn) UpdateData(id int64, nombre, apellido, matricula string) error {
	var myOdontologo odontologoModel.Odontologos
	errUpdate := config.DbConn.Model(&myOdontologo).Where("id = ?", id).Updates(odontologoModel.Odontologos{
		Nombre:        nombre,
		Apellido: apellido,
		Matricula: matricula,
	}).Error

	if errUpdate != nil {
		return errors.New("UpdateData errUpdate = " + errUpdate.Error())
	}

	return nil
}

// DeleteData - soft delete data in table odontologo_s
func (config *dbConn) DeleteData(id int64) error {
	var myOdontologo odontologoModel.Odontologos
	errDelete := config.DbConn.Where("id = ?", id).Delete(&myOdontologo).Error
	if errDelete != nil {
		return errors.New("DeleteData errDelete = " + errDelete.Error())
	}

	return nil
}

// FindAll - get all data in table odontologo_s
func (config *dbConn) FindAll() ([]odontologoModel.Odontologos, error) {
	var list []odontologoModel.Odontologos

	err := config.DbConn.Select("id,apellido, nombre, matricula").Find(&list).Error
	if err != nil {
		return list, errors.New("FindAll err = " + err.Error())
	}

	return list, nil
}

// FindByID - get data odontologo_s by id
func (config *dbConn) FindByID(id int64) (odontologoModel.Odontologos, error) {
	var data odontologoModel.Odontologos
	err := config.DbConn.Select("id,nombre,apellido,matricula").Where("id = ?", id).First(&data).Error
	if err != nil {
		if err.Error() == "record not found" {
			return data, nil
		}
		return data, errors.New("FindByID err = " + err.Error())
	}

	return data, nil
}