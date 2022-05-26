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

type ExperienceRepository struct {
	db *gorm.DB
}

func NewExperienceRepo() *ExperienceRepository {
	db := config.SetupDB()

	return &ExperienceRepository{db: db}
}

func (repository *ExperienceRepository) CreateExperience(c *gin.Context) {
	var experience models.Experience
	c.BindJSON(&experience)
	err := models.CreateExperience(repository.db, &experience)
	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), experience)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "Success", experience)
	c.JSON(http.StatusOK, res)
}

//get experiences
func (repository *ExperienceRepository) ReadExperience(c *gin.Context) {
	var experience []models.Experience
	id, _ := c.Params.Get("id")
	idInt, _ := strconv.ParseUint(id, 0, 64)

	err := models.ReadExperienceByUser(repository.db, &experience, idInt)

	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), experience)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "Success", experience)
	c.JSON(http.StatusOK, res)
}

//get experience by id
func (repository *ExperienceRepository) ReadExperienceById(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var experience models.Experience
	idInt, _ := strconv.ParseUint(id, 0, 64)
	err := models.ReadExperienceById(repository.db, &experience, idInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res := helper.BuildErrorResponse("Record not found", err.Error(), experience)
			c.AbortWithStatusJSON(http.StatusNotFound, res)
			return
		}

		res := helper.BuildErrorResponse("Error", err.Error(), experience)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "Success", experience)
	c.JSON(http.StatusOK, res)
}

// update experience
func (repository *ExperienceRepository) UpdateExperience(c *gin.Context) {
	var experience models.Experience
	id, _ := c.Params.Get("id")
	idInt, _ := strconv.ParseUint(id, 0, 64)

	err := models.ReadExperienceById(repository.db, &experience, idInt)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res := helper.BuildErrorResponse("Record not found", err.Error(), experience)
			c.AbortWithStatusJSON(http.StatusNotFound, res)
			return
		}

		res := helper.BuildErrorResponse("Error", err.Error(), experience)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	c.BindJSON(&experience)

	err = models.UpdateExperience(repository.db, &experience)

	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), experience)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "Experience updated successfully", experience)
	c.JSON(http.StatusOK, res)
}

// delete experience
func (repository *ExperienceRepository) DeleteExperience(c *gin.Context) {
	var experience models.Experience

	id, _ := c.Params.Get("id")
	idInt, _ := strconv.ParseUint(id, 0, 64)

	err := models.DeleteExperience(repository.db, &experience, idInt)

	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), experience)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "Experience deleted  successfully", experience)
	c.JSON(http.StatusOK, res)
}
