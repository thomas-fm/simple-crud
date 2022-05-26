package repository

import (
	"test-yukbisnis/models"

	"gorm.io/gorm"
)

type CVRepository interface {
	CreateCV(cv models.CV) models.CV
	UpdateCV(cv models.CV) models.CV
	ReadCV() []models.CV
	ReadCVById(cv_id uint64) models.CV
	DeleteCV(cv models.CV)
}

type cvConnection struct {
	connection *gorm.DB
}

func NewCVRepository(connection *gorm.DB) CVRepository {
	return &cvConnection{
		connection: connection,
	}
}

func (db *cvConnection) CreateCV(cv models.CV) models.CV {
	db.connection.Save(&cv)
	db.connection.Find(&cv)
	return cv
}

// READ
func (db *cvConnection) ReadCV() []models.CV {
	var cvs []models.CV
	db.connection.Find(&cvs)

	return cvs
}

// UPDATE
func (db *cvConnection) UpdateCV(cv models.CV) models.CV {
	db.connection.Save(&cv)
	db.connection.Find(&cv)
	return cv
}

// DELETE
func (db *cvConnection) DeleteCV(cv models.CV) {
	db.connection.Delete(&cv)
}

// find by id
func (db *cvConnection) ReadCVById(cvID uint64) models.CV {
	var cv models.CV
	db.connection.Find(&cv, cvID)
	return cv
}
