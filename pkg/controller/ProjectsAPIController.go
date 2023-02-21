package controller

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ssjh23/PortfolioBackend/pkg/common/models"
	"github.com/ssjh23/PortfolioBackend/pkg/service"
	"gorm.io/gorm"
)

type ProjectsAPIController struct {
	srv *service.ProjectsService
}

// Create address of APIService layer
func NewProjectsAPIController(db *gorm.DB) *ProjectsAPIController {
	return &ProjectsAPIController{
		srv: service.NewProjectsService(db),
	}
}
// Pointer receiver to DB


// POST Request Add Project
func (api ProjectsAPIController) AddProject(c *gin.Context){
	err := c.Request.ParseMultipartForm(12000)
	if err != nil {
		fmt.Println("Error Retrieving Form Body")
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	var req models.ProjectForm
	req.Title = c.Request.Form.Get("title")
	req.Description = c.Request.Form.Get("description")
	req.Techstack= c.Request.Form.Get("techstack")
	resp := api.srv.AddProject(&req)
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Success",
			"data": resp,
		},
	)
}

func (api ProjectsAPIController) GetProjects (c *gin.Context){
	getProjectQueries := models.GetProjectQueries{}
	UrlQueries := c.Request.URL.Query()
	stringCount := UrlQueries.Get("count")
	stringSortedBy := UrlQueries.Get("sortBy")
	stringOffset := UrlQueries.Get("offset")
	count,_ := strconv.Atoi(stringCount)
	offset,_ := strconv.Atoi(stringOffset)
	getProjectQueries.Count = count
	getProjectQueries.SortedBy = stringSortedBy
	getProjectQueries.Offset = offset
	resp := api.srv.GetProject(&getProjectQueries)
	code := http.StatusOK
	message := "Success"
	// The error here is a valid query param in terms of the input domain, but limited by the data present in the "database"
	if (resp.Error != ""){
		code = http.StatusBadRequest
		message = "Query Error"
	}
	c.JSON(
		code,
		gin.H{
			"status": code,
			"message": message,
			"data": resp,
		},
	)
}

func (api ProjectsAPIController) DeleteProject (c *gin.Context){
	UrlQueries := c.Request.URL.Query()
	title := UrlQueries.Get("title")
	queryBody := models.DeleteProjectQuery{
		Title: title,
	}
	resp := api.srv.DeleteProject(&queryBody)
	code := http.StatusOK
	message := "Success"
	if (resp.Error != ""){
		code = http.StatusBadRequest
		message = "Query Error"
	}
	c.JSON(
		code,
		gin.H{
			"status": code,
			"message": message,
			"data": resp,
		},
	)
	
}
