package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/ssjh23/PortfolioBackend/pkg/common/db"
	"github.com/ssjh23/PortfolioBackend/pkg/router"
)

func main() {
	r := gin.Default()
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()
    router.Routes(r)
    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)
    fmt.Println("Running on port" + port)
    db.Init(dbUrl)
    r.Run(port)
}