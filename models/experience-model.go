package models

import (
	"gorm.io/gorm"
)

func CreateExperience(db *gorm.DB, experience *Experience) (err error) {
	err = db.Create(experience).Error
	if err != nil {
		return err
	}

	return nil
}

func ReadExperience(db *gorm.DB, experiences *[]Experience) (err error) {
	err = db.Find(&experiences).Error

	if err != nil {
		return err
	}

	return nil
}

func ReadExperienceById(db *gorm.DB, experience *Experience, id uint64) (err error) {
	err = db.Where("id = ?", id).First(experience).Error

	if err != nil {
		return err
	}

	return nil
}

func ReadExperienceByUser(db *gorm.DB, experience *[]Experience, cv_id uint64) (err error) {
	err = db.Where("cv_id = ?", cv_id).Find(experience).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateExperience(db *gorm.DB, experience *Experience) (err error) {
	err = db.Save(experience).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteExperience(db *gorm.DB, experience *Experience, id uint64) (err error) {
	err = db.Where("id = ?", id).Delete(experience).Error
	if err != nil {
		return err
	}
	return nil
}
