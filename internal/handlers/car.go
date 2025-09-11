package handlers

import (
	"Webserver/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}

	err := c.BindJSON(&body)
	if err != nil {
		return
	}

	car, _ := services.Create(c, body.Name)

	c.JSON(201, gin.H{"data": car})
}

func GetAll(c *gin.Context) {
	cars, _ := services.GetAll(c)

	c.JSON(200, gin.H{"data": cars})
}

func Update(c *gin.Context) {
	idRaw := c.Param("id")

	id, _ := strconv.Atoi(idRaw)

	var body struct {
		Name  *string `json:"name,omitempty"`
		Price *int    `json:"price,omitempty"`
	}

	err := c.BindJSON(&body)
	if err != nil {
		return
	}

	car, _ := services.Update(c, id, body.Name, body.Price)

	c.JSON(201, gin.H{"data": car})
}

func SetYear(c *gin.Context) {
	idRaw := c.Param("id")

	id, _ := strconv.Atoi(idRaw)

	var body struct {
		Year int `json:"year"`
	}

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	car, _ := services.SetYear(c, id, body.Year)

	c.JSON(201, gin.H{"data": car})
}

func Delete(c *gin.Context) {
	idRaw := c.Param("id")

	id, _ := strconv.Atoi(idRaw)
	car, _ := services.Delete(c, id)

	c.JSON(201, gin.H{"data": car})
}
