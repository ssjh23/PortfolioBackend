package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CheckAddProjectReqBody (c *gin.Context){
	err := c.Request.ParseMultipartForm(12000)
	if err != nil {
		fmt.Println("Error Retrieving Form Body")
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	title := c.Request.Form.Get("title")
	description := c.Request.Form.Get("description")
	techstack := c.Request.Form.Get("techstack")
	 
	error := []string{}
	if (title == ""){
		error = append(error, "Title is empty")
	}
	if (description == ""){
		error = append(error, "Description is empty")
	}
	if (techstack == ""){
		error = append(error, "Tech stack is empty")
	}
	if (len(error)!= 0){
		c.JSON(
			http.StatusBadRequest, 
			gin.H{
				"Error": error,
		})
		c.Abort()
	}
	return
}

func CheckGetProjectsQueries(c *gin.Context){
	UrlQueries := c.Request.URL.Query()
	stringCount := UrlQueries.Get("count")
	stringSortedBy := UrlQueries.Get("sortBy")
	stringOffset := UrlQueries.Get("offset")
	error := []string{}
	if (stringCount != ""){
		_, err := strconv.Atoi(stringCount)
		if err != nil {
			error = append(error, "Invalid Count Query")
			c.Error(err).SetType(gin.ErrorTypePublic)

		}
	}
	if (stringSortedBy != ""){
		if (stringSortedBy != "Title"){
			error = append(error, "Invalid sortBy Query")
			
		}
	}
	if (stringOffset != ""){
		_,err := strconv.Atoi(stringOffset)
		if err != nil {
			error = append(error, "Invalid Offseet Query")
			
		}
	}
	if (len(error)!= 0){
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Error": error,
			},
		)
		c.Abort()
	}
	return

}