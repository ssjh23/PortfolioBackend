package controller

import (
	"net/http"
	"io/ioutil"

	"github.com/ssjh23/PortfolioBackend/pkg/service"
    "github.com/gin-gonic/gin"

	// "gorm.io/gorm"
)

type AboutAPIController struct {
	srv *service.AboutService
}

// Create address of APIService layer
func NewAboutAPIController() *AboutAPIController {
	return &AboutAPIController{
		srv: service.NewAboutService(),
	}
}
// Pointer receiver to DB


// POST Request Hero Description
func (AboutAPIController) GetProfilePicture(c *gin.Context){
	fileBytes, err := ioutil.ReadFile("reduce.png")
	if err != nil {
		panic(err)
	}
	
	c.Writer.WriteHeader(http.StatusAccepted)
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Success",
			"data":    fileBytes,
		},
	)
}