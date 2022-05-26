package models

import (
	"gorm.io/gorm"
)

func CreateCV(db *gorm.DB, cv *CV) (err error) {
	err = db.Create(cv).Error
	if err != nil {
		return err
	}

	return nil
}

func ReadCV(db *gorm.DB, cvs *[]CV) (err error) {
	err = db.Preload("Experiences").Preload("Educations").Preload("Contacts").Find(&cvs).Error

	if err != nil {
		return err
	}

	return nil
}

func ReadCVById(db *gorm.DB, cv *CV, id uint64) (err error) {
	err = db.Where("id = ?", id).First(cv).Error

	if err != nil {
		return err
	}

	db.Model(&cv).Association("Educations").Find(&cv.Educations)
	db.Model(&cv).Association("Experiences").Find(&cv.Experiences)
	db.Model(&cv).Association("Contacts").Find(&cv.Contacts)

	return nil
}

func UpdateCV(db *gorm.DB, cv *CV) (err error) {
	err = db.Save(cv).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteCV(db *gorm.DB, cv *CV, id uint64) (err error) {
	err = db.Where("id = ?", id).Delete(cv).Error
	if err != nil {
		return err
	}
	return nil
}
