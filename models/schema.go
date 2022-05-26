package models

import "gorm.io/gorm"

type CV struct {
	gorm.Model
	Name        string       `json:"name"`
	Age         int          `json:"age"`
	Summary     string       `json:"summary"`
	Photo       string       `json:"photo"`
	Educations  []Education  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"educations"`
	Experiences []Experience `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"  json:"experiences"`
	Contacts    []Contact    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"  json:"contacts"`
}

type Education struct {
	gorm.Model
	CVID        int    `json:"cv_id"`
	Title       string `json:"title"`
	Major       string `json:"major"`
	School      string `json:"school"`
	Description string `json:"description"`
	StartYear   int    `json:"start_year"`
	EndYear     int    `json:"end_year"`
}

type Experience struct {
	gorm.Model
	CVID        int    `json:"cv_id"`
	Title       string `json:"title"`
	Company     string `json:"company"`
	Description string `json:"description"`
	StartYear   int    `json:"start_year"`
	EndYear     int    `json:"end_year"`
}

type Contact struct {
	gorm.Model
	CVID  int    `json:"cv_id"`
	Type  string `json:"type"`
	Value string `json:"value"`
}
