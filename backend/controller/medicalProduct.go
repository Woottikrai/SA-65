package controller

import (
	"github.com/cottoncandyblue/sa-65-project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /medical_product

func CreateMedicalProduct(c *gin.Context) {

	var mp entity.MedicalProduct

	if err := c.ShouldBindJSON(&mp); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&mp).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": mp})

}

// GET /medical_products

func ListMedicalProduct(c *gin.Context) {

	var mps []entity.MedicalProduct
	if err := entity.DB().Raw("SELECT * FROM medical_products").Find(&mps).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mps})

}
