package main

import (
	"go_app/user"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	router.GET("/", handler)
	router.GET("/env", func(c *gin.Context) {
		name := os.Getenv("NAME")
		c.String(http.StatusOK, "Hello "+name)
	})
	router.Run()
}

func handler(c *gin.Context) {
	dsn := "swd:swd_password@tcp(mariadb:3306)/goapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)

}
