package main

import (
	"exercise_go/src"
	"exercise_go/src/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("SQL_URI").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	src.RegisterRoutes(r, h)

	r.Run(port)
}
