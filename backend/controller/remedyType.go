package controller

import (
	"github.com/cottoncandyblue/sa-65-project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /remedy_type

func CreateRemedyType(c *gin.Context) {

	var RemedyType entity.RemedyType

	if err := c.ShouldBindJSON(&RemedyType); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&RemedyType).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": RemedyType})

}

func GetRemedyType(c *gin.Context) {
	var RemedyType entity.RemedyType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM remedy_types WHERE id = ?", id).Scan(&RemedyType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": RemedyType})
}

// GET /remedy_types

func ListRemedyType(c *gin.Context) {

	var remedytypes []entity.RemedyType
	if err := entity.DB().Raw("SELECT * FROM remedy_types").Find(&remedytypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": remedytypes})

}
