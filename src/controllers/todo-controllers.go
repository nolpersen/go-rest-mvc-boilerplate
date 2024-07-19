package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nolpersen/src/config"
	"github.com/nolpersen/src/models"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB()

type todoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type todoResponse struct {
	todoRequest
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ID        uint   `json:"id"`
}

func Index(context *gin.Context) {
	var todos []models.Todo

	// Querying to find todo datas.
	err := db.Find(&todos)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todos,
	})
}

func Show() {

}

func Store(context *gin.Context) {
	var data todoRequest

	//binding req body json to struct
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{}
	todo.Name = data.Name
	todo.Description = data.Description
	todo.Status = data.Status
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	result := db.Create(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	// Matching result to create response
	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description
	response.Status = todo.Status
	response.CreatedAt = todo.CreatedAt.GoString()
	response.UpdatedAt = todo.UpdatedAt.GoString()

	// Creating http response
	context.JSON(http.StatusCreated, response)

}

func Update(context *gin.Context) {
	var data todoRequest

	// Defining request parameter to get todo id
	reqParamId := context.Param("id")
	id := cast.ToUint(reqParamId)

	// Binding request body json to request body struct
	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Initiate models todo
	todo := models.Todo{}

	// Querying find todo data by todo id from request parameter
	todoById := db.Where("id = ?", id).First(&todo)
	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	// Matching todo request with todo models
	todo.Name = data.Name
	todo.Description = data.Description

	// Update new todo data
	result := db.Save(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	// Matching result to todo response struct
	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	// Creating http response
	context.JSON(http.StatusCreated, response)
}

func Delete(context *gin.Context) {
	// Initiate todo models
	todo := models.Todo{}
	// Getting request parameter id
	reqParamId := context.Param("id")
	id := cast.ToUint(reqParamId)

	// Querying delete todo by id
	delete := db.Where("id = ?", id).Unscoped().Delete(&todo)
	fmt.Println(delete)

	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    id,
	})

}
