package models

type GetProjectQueries struct {
	SortedBy string
	Count int
	Offset int
}

type ProjectForm struct {
	Title string `form:"title"`
	Description string `form:"description"`
	Techstack string `form:"techstack"`
}

type ProjectItem struct {
	Title string
	Description string
	Techstack []string
}

type AddProjectResp struct {
	Projects []ProjectItem
}

type DeleteProjectResp struct {
	Error string
	Projects []ProjectItem
}

type GetProjectsResp struct {
	Error string
	Queries  GetProjectQueries
	Projects []ProjectItem
}

type DeleteProjectQuery struct {
	Title string 
}