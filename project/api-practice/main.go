package main

import (
	"api-practice/config"
	"api-practice/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": "masuk!",
		})
	})

	router.POST("/mahasiswa", inDB.CreateMahasiswa)
	router.GET("/mahasiswas", inDB.GetMahasiswas)
	router.GET("/mahasiswa/:id", inDB.GetMahasiswa)
	router.PUT("/mahasiswa/:id", inDB.UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id", inDB.DeleteMahasiswa)

	router.Run(":3000")
}
