package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Routes (r *gin.Engine){
    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)
	route := r.Group("")
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
		HeroDescription := route.Group("/HeroDesc")
		{
			HeroDescription.GET("/")
		}
	}
}