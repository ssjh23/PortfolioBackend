package project

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ssjh23/PortfolioBackend/pkg/common/models"
)

type ProjectRepo struct {}

var project1 = models.ProjectItem{
	Title: "Scancart",
	Description: "An Android App tailored to aiding the needs of the Elderly",
	Techstack: []string{"Android Studio", "Java", "Firebase"},
}

var project2 = models.ProjectItem{
	Title: "What the Hack 2021",
	Description: "A user-friendly IOT system that aims to boost the efficiencies of the waste-collection process in Singapore by reducing the travelling time and workload of waste-management employees",
	Techstack: []string{"Python","Telegram API","Firebase","Google Maps API"},
}
var existing_projects = []models.ProjectItem{
	project1,
	project2,
}
var existing_proj_pointer = &existing_projects


func (repo ProjectRepo) AddProject(req *models.ProjectForm)(resp models.AddProjectResp){
	techStackList := strings.Split(req.Techstack, ",")
	newProject := models.ProjectItem{
		Title: req.Title,
		Description: req.Description,
		Techstack: techStackList,
	}
	s := append((*existing_proj_pointer),newProject)
	existing_proj_pointer = &s
	resp = models.AddProjectResp{Projects: (*existing_proj_pointer)}
	fmt.Println(resp)
	return resp
}

// The order to sort the queries are 1. SortedBy 2. Offset 3. Count
func (repo ProjectRepo) GetProjects(queries *models.GetProjectQueries)(resp models.GetProjectsResp){
	sortedBy := queries.SortedBy
	offset := queries.Offset
	count := queries.Count

	resp = models.GetProjectsResp{
		Queries: (*queries),
	}

	allProjects := (*existing_proj_pointer)
	numberOfProjects := len(allProjects)
	// Sorts the list alphabetically using title
	if (sortedBy != ""){
	sort.Slice(allProjects, func(i, j int) bool {
			return allProjects[i].Title < allProjects[j].Title
		})
	}
	// If the offset exceeds the number of projects, return error
	if (offset >= numberOfProjects){
		resp.Error = "The offset exceeds the number of project. Number of projects is " + strconv.Itoa(numberOfProjects)
		return resp
	}
	// If the count in the query is more than the number of projects or 0, return all the projects
	if (count == 0){
		count = len(allProjects)
	}
	if (count > numberOfProjects){
		resp.Error = "The requested count exceeds the number of projects. Number of projects: "  + strconv.Itoa(numberOfProjects)
		return resp
	}
	if (count+offset > numberOfProjects){
		resp.Error = "Cannot return stated number of projects due to offset. Number of projects: "  + strconv.Itoa(numberOfProjects)
		return resp
	}
	// use slice index to get a portion of the array of projects
	slicedProjects := allProjects[offset:count+offset]
	resp.Projects = slicedProjects
	return resp
}