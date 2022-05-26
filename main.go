package main

import (
	"net/http"
	"test-yukbisnis/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello world")
	})

	cvController := controllers.New()
	eduController := controllers.NewEducationRepo()
	expController := controllers.NewExperienceRepo()

	r.POST("/cv", cvController.CreateCV)
	r.GET("/cv", cvController.ReadCV)
	r.GET("/cv/:id", cvController.ReadCVById)
	r.PUT("/cv/:id", cvController.UpdateCV)
	r.DELETE("/cv/:id", cvController.DeleteCV)

	r.POST("/education", eduController.CreateEducation)
	r.GET("/education/:id/user", eduController.ReadEducation)
	r.GET("/education/:id", eduController.ReadEducationById)
	r.PUT("/education/:id", eduController.UpdateEducation)
	r.DELETE("/education/:id", eduController.DeleteEducation)

	r.POST("/experience", expController.CreateExperience)
	r.GET("/experience/:id/user", expController.ReadExperience)
	r.GET("/experience/:id", expController.ReadExperienceById)
	r.PUT("/experience/:id", expController.UpdateExperience)
	r.DELETE("/experience/:id", expController.DeleteExperience)

	r.POST("/upload-photo", cvController.UploadPhoto)

	r.StaticFS("/file", http.Dir("public"))

	return r
}
