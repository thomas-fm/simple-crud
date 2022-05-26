package controllers

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"test-yukbisnis/config"
	"test-yukbisnis/helper"
	"test-yukbisnis/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CVRepository struct {
	db *gorm.DB
}

func New() *CVRepository {
	db := config.SetupDB()

	return &CVRepository{db: db}
}

func (repository *CVRepository) CreateCV(c *gin.Context) {
	var cv models.CV
	c.BindJSON(&cv)
	err := models.CreateCV(repository.db, &cv)
	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), cv)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "Success", cv)
	c.JSON(http.StatusOK, res)
}

//get cvs
func (repository *CVRepository) ReadCV(c *gin.Context) {
	var cv []models.CV
	err := models.ReadCV(repository.db, &cv)
	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), cv)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "Success", cv)
	c.JSON(http.StatusOK, res)
}

//get cv by id
func (repository *CVRepository) ReadCVById(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var cv models.CV
	idInt, _ := strconv.ParseUint(id, 0, 64)
	err := models.ReadCVById(repository.db, &cv, idInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res := helper.BuildErrorResponse("Record not found", err.Error(), cv)
			c.AbortWithStatusJSON(http.StatusNotFound, res)
			return
		}

		res := helper.BuildErrorResponse("Error", err.Error(), cv)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "Success", cv)
	c.JSON(http.StatusOK, res)
}

// update cv
func (repository *CVRepository) UpdateCV(c *gin.Context) {
	var cv models.CV
	id, _ := c.Params.Get("id")
	idInt, _ := strconv.ParseUint(id, 0, 64)

	err := models.ReadCVById(repository.db, &cv, idInt)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res := helper.BuildErrorResponse("Record not found", err.Error(), cv)
			c.AbortWithStatusJSON(http.StatusNotFound, res)
			return
		}

		res := helper.BuildErrorResponse("Error", err.Error(), cv)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	c.BindJSON(&cv)

	err = models.UpdateCV(repository.db, &cv)

	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), cv)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "CV updated successfully", cv)
	c.JSON(http.StatusOK, res)
}

// delete cv
func (repository *CVRepository) DeleteCV(c *gin.Context) {
	var cv models.CV

	id, _ := c.Params.Get("id")
	idInt, _ := strconv.ParseUint(id, 0, 64)

	err := models.DeleteCV(repository.db, &cv, idInt)

	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), cv)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "CV deleted  successfully", cv)
	c.JSON(http.StatusOK, res)
}

func (repository *CVRepository) UploadPhoto(c *gin.Context) {
	var cv models.CV

	// Check if id exist
	id := c.PostForm("cv_id")
	idInt, _ := strconv.ParseUint(id, 0, 64)
	err := models.ReadCVById(repository.db, &cv, idInt)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res := helper.BuildErrorResponse("Record not found", err.Error(), cv)
			c.AbortWithStatusJSON(http.StatusNotFound, res)
			return
		}

		res := helper.BuildErrorResponse("Error", err.Error(), cv)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	// Upload file
	file, header, err := c.Request.FormFile("file")

	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), cv)
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	filename := header.Filename
	out, err := os.Create("public/" + filename)

	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()
	_, err = io.Copy(out, file)

	if err != nil {
		log.Fatal(err)
	}

	filepath := "http://localhost:8080/file/" + filename
	cv.Photo = filepath
	c.BindJSON(&cv)

	err = models.UpdateCV(repository.db, &cv)

	if err != nil {
		res := helper.BuildErrorResponse("Error", err.Error(), cv)
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "Photo updated successfully", cv)
	c.JSON(http.StatusOK, res)
}
