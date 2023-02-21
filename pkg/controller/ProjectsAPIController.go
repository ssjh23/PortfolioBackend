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
	req_pointer := &req
	req.Title = c.Request.Form.Get("title")
	req.Description = c.Request.Form.Get("description")
	req.Techstack= c.Request.Form.Get("techstack")
	resp := api.srv.AddProject(req_pointer)
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
	getProjectQueriesPointer := &getProjectQueries
	UrlQueries := c.Request.URL.Query()
	stringCount := UrlQueries.Get("count")
	stringSortedBy := UrlQueries.Get("sortBy")
	stringOffset := UrlQueries.Get("offset")
	if (stringCount != ""){
		count,err := strconv.Atoi(stringCount)
		if err != nil {
			error := "Invalid Count Query"
			fmt.Println(error)
			c.Error(err).SetType(gin.ErrorTypePublic)
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"Error": error,
				},
			)
			return
		}
		getProjectQueriesPointer.Count = count
	}
	if (stringSortedBy != ""){
		if (stringSortedBy != "Title"){
			error := "Invalid sortedBy Query"
			fmt.Println(error)
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"Error": error,
				},
			)
			return
		}
		getProjectQueriesPointer.SortedBy = stringSortedBy
	}
	if (stringOffset != ""){
		offset,err := strconv.Atoi(stringOffset)
		if err != nil {
			error := "Invalid Offset Query"
			fmt.Println(error)
			c.Error(err).SetType(gin.ErrorTypePublic)
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"Error": error,
				},
			)
			return
		}
		getProjectQueriesPointer.Offset = offset
	}
	resp := api.srv.GetProject(getProjectQueriesPointer)
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
	fmt.Println(resp)
}
