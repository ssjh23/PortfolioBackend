package service

import (

	"github.com/ssjh23/PortfolioBackend/pkg/common/models"
	"github.com/ssjh23/PortfolioBackend/pkg/repository"
	// "gorm.io/gorm"
)

type ProjectsService struct {
	// db       *gorm.DB
	projects *project.ProjectRepo
}

func NewProjectsService() *ProjectsService {
	return &ProjectsService{
		// db:       db,
		projects: &project.ProjectRepo{},
	}
}

func (srv ProjectsService) AddProject(req *models.ProjectForm)(resp models.AddProjectResp){
	resp = srv.projects.AddProject(req)
	return
}

func (srv ProjectsService) GetProject(queries *models.GetProjectQueries)(resp models.GetProjectsResp){
	resp = srv.projects.GetProjects(queries)
	return
}

func (srv ProjectsService) DeleteProject(query *models.DeleteProjectQuery)(resp models.DeleteProjectResp){
	resp = srv.projects.DeleteProject(query)
	return
}