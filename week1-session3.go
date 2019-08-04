package main

import (
	"time"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id int
	Title string
	Completed bool
	CreatedAt time.Time
}

var db *gorm.DB

func runWeek1Session3() {
	var err error
	db, err = gorm.Open("mysql", "root@tcp(localhost:3306)/sys?parseTime=True")
	if err != nil {
		log.Fatal(err)
		log.Fatal("Failed to connection DB")
	}
	defer db.Close()

	db.LogMode(true)
	err = db.AutoMigrate(Todo{}).Error
	if err != nil {
		log.Fatal("failed to migrate db table todo")
	}

	router := gin.Default()

	router.GET("/todos", listTodos)
	router.POST("/create", createTodo)
	router.Run(":8008")

	return
}

func listTodos(c *gin.Context) {
	var todos []Todo
	err := db.Find(&todos).Error

	if err != nil {
		log.Fatal(err)
		c.String(500, "Failed to list Todo")
		return
	}
	
	c.JSON(200, todos)
}



func createTodo(c *gin.Context) {
	var argument struct {
		Title string
	}

	err := c.BindJSON(&argument)
	if err != nil {
		c.String(400, "invalid param")
		return
	}

	todo := Todo {
		Title: argument.Title,
	}

	err = db.Create(&todo).Error
	if err != nil {
		c.String(500, "failed to create new todo")
		return
	}

	c.JSON(200, todo)
}

