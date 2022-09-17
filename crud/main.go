package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// DBスキーマ
type Todo struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name      string         `json:"name"`
}

func dbConnect() *gorm.DB {
	dsn := "xxx:xxx@tcp(mysql:3306)/todo?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func main() {
	engine := gin.Default()
	db := dbConnect()
	db.AutoMigrate(&Todo{})

	// create
	engine.POST("/todo", func(c *gin.Context) {
		todo := Todo{}
		c.BindJSON(&todo)
		if err := db.Create(&todo).Error; err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, todo)
	})

	// read
	engine.GET("/todo/:id", func(c *gin.Context) {
		todo := Todo{}
		id := c.Param("id")
		if err := db.First(&todo, id).Error; err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, todo)
	})

	// update
	engine.PUT("/todo/:id", func(c *gin.Context) {
		todo := Todo{}
		id := c.Param("id")
		if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
			fmt.Println(err)
		}
		c.BindJSON(&todo)
		db.Save(&todo)
		c.JSON(http.StatusOK, todo)
	})

	// delete
	engine.DELETE("/todo/:id", func(c *gin.Context) {
		todo := Todo{}
		id := c.Param("id")
		if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
			fmt.Println(err)
		}
		db.Delete(&todo)
		c.JSON(http.StatusOK, todo)
	})

	engine.Run(":8085")
}