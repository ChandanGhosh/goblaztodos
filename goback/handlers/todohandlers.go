package handlers

import (
	"encoding/json"
	"github.com/chandanghosh/goblaztodos/goback/todo"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

// GetAllTodoHandler ...
func GetAllTodoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, todo.GetAllTodo())
}

// AddTodoHandler ...
func AddTodoHandler(c *gin.Context) {
	content, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error occured: %v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var todoItem todo.Todo
	err = json.Unmarshal(content, &todoItem)
	if err != nil {
		log.Printf("Error in unmarshaling content, %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": todo.AddTodo(todoItem.Message)})
}

// UpdateTodoHandler ...
func UpdateTodoHandler(c *gin.Context) {
	content, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer c.Request.Body.Close()
	var todoItem todo.Todo
	err = json.Unmarshal(content, &todoItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if err = todo.UpdateToDo(todoItem); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	} 
	
	c.JSON(http.StatusOK, "")

}
