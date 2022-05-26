package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"test-yukbisnis/config"
	"test-yukbisnis/helper"
	"test-yukbisnis/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EducationRepository struct {
	db *gorm.DB
}

func NewEducationRepo() *EducationRepository {
	db := config.SetupDB()

	return &EducationRepository{db: db}
}

func (repository *EducationRepository) CreateEducation(c *gin.Context) {
	var education models.Education
	c.BindJSON(&education)
	err := models.CreateEducation(repository.db, &education)
	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), education)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "Success", education)
	c.JSON(http.StatusOK, res)
}

//get educations
func (repository *EducationRepository) ReadEducation(c *gin.Context) {
	var education []models.Education

	id, _ := c.Params.Get("id")
	idInt, _ := strconv.ParseUint(id, 0, 64)

	err := models.ReadEducationByUser(repository.db, &education, idInt)
	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), education)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "Success", education)
	c.JSON(http.StatusOK, res)
}

//get education by id
func (repository *EducationRepository) ReadEducationById(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var education models.Education
	idInt, _ := strconv.ParseUint(id, 0, 64)
	err := models.ReadEducationById(repository.db, &education, idInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res := helper.BuildErrorResponse("Record not found", err.Error(), education)
			c.AbortWithStatusJSON(http.StatusNotFound, res)
			return
		}

		res := helper.BuildErrorResponse("Error", err.Error(), education)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "Success", education)
	c.JSON(http.StatusOK, res)
}

// update education
func (repository *EducationRepository) UpdateEducation(c *gin.Context) {
	var education models.Education
	id, _ := c.Params.Get("id")
	idInt, _ := strconv.ParseUint(id, 0, 64)

	err := models.ReadEducationById(repository.db, &education, idInt)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res := helper.BuildErrorResponse("Record not found", err.Error(), education)
			c.AbortWithStatusJSON(http.StatusNotFound, res)
			return
		}

		res := helper.BuildErrorResponse("Error", err.Error(), education)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	c.BindJSON(&education)

	err = models.UpdateEducation(repository.db, &education)

	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), education)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "Education updated successfully", education)
	c.JSON(http.StatusOK, res)
}

// delete education
func (repository *EducationRepository) DeleteEducation(c *gin.Context) {
	var education models.Education

	id, _ := c.Params.Get("id")
	idInt, _ := strconv.ParseUint(id, 0, 64)

	err := models.DeleteEducation(repository.db, &education, idInt)

	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), education)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "Education deleted  successfully", education)
	c.JSON(http.StatusOK, res)
}
