package models

import (
	"gorm.io/gorm"
)

func CreateEducation(db *gorm.DB, education *Education) (err error) {
	err = db.Create(education).Error
	if err != nil {
		return err
	}

	return nil
}

func ReadEducation(db *gorm.DB, educations *[]Education) (err error) {
	err = db.Find(&educations).Error

	if err != nil {
		return err
	}

	return nil
}

func ReadEducationById(db *gorm.DB, education *Education, id uint64) (err error) {
	err = db.Where("id = ?", id).First(education).Error

	if err != nil {
		return err
	}

	return nil
}

func ReadEducationByUser(db *gorm.DB, education *[]Education, cv_id uint64) (err error) {
	err = db.Where("cv_id = ?", cv_id).Find(education).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateEducation(db *gorm.DB, education *Education) (err error) {
	err = db.Save(education).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteEducation(db *gorm.DB, education *Education, id uint64) (err error) {
	err = db.Where("id = ?", id).Delete(education).Error
	if err != nil {
		return err
	}
	return nil
}
