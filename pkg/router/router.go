package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	api "github.com/ssjh23/PortfolioBackend/pkg/controller"
	"github.com/ssjh23/PortfolioBackend/pkg/common/db"
)


func Routes (r *gin.Engine){
    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)
	route := r.Group("")
	orm := db.Init((dbUrl))
	{
		route.GET("/ping",func(ctx *gin.Context){
			ctx.JSON(
				http.StatusOK,
				gin.H{
					"status":  http.StatusOK,
					"message": "Pong!",
					"data":    nil,
					"port": port,
					"dbUrl": dbUrl,
				},
			)
		})
		AboutAPI := api.NewAboutAPIController(orm)
		About := route.Group("/about")
		{
			About.GET("/get_profile_pic",AboutAPI.GetProfilePicture)
		}
		ProjectsAPI := api.NewProjectsAPIController(orm)
		Projects := route.Group("/project")
		{
			Projects.POST("/add_project", ProjectsAPI.AddProject)
			Projects.GET("/get_projects",ProjectsAPI.GetProjects)
		}

	}
}