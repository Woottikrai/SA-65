/***********************************
 *	THIS IS UNTESTED CONTROLLER    *
 ***********************************/

 package controller

 import (
	 "net/http"
 
	 "github.com/gin-gonic/gin"
	 "github.com/cottoncandyblue/sa-65-project/entity"
 )
 
//List /treatmentRecord
func ListTreatment(c *gin.Context) {
	var treatmentRecord []entity.Treatment
	if err := entity.DB().Preload("Screening").
		Preload("Screening.Patient").
		Table("treatments").
		Find(&treatmentRecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": treatmentRecord})
}

// POST /treatmentRecord
func CreateTreatment(context *gin.Context) {
	var treatmentRecord entity.Treatment

	var screening entity.Screening
	var userdentist entity.User
	var remedy entity.RemedyType

	if err := context.ShouldBindJSON(&treatmentRecord); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", treatmentRecord.UserDentistID).First(&userdentist); tx.RowsAffected == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Dentist not found"})
		return
	}

	entity.DB().Joins("Role").Find(&userdentist)

	if userdentist.Role.Name != "Dentist" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "only for userdentsit"})
		return
	}

	if tx := entity.DB().Where("id = ?", treatmentRecord.ScreeningID).First(&screening); tx.RowsAffected == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Screening not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", treatmentRecord.RemedyTypeID).First(&remedy); tx.RowsAffected == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "RemedyType not found"})
		return
	}

	treatmentData := entity.Treatment{
		PrescriptionRaw:  treatmentRecord.PrescriptionRaw,
		ToothNumber:      treatmentRecord.ToothNumber,	
		// create with assosiation
		Screening:   screening,
		UserDentist: userdentist,
		RemedyType:  remedy,
	}

	if err := entity.DB().Create(&treatmentData).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": treatmentData})
}
// GET /treatment
func ListTreatments(c *gin.Context) {
	var treatments []entity.Treatment

	// preload association
	if err := entity.DB().Preload("Screening.Patient").Preload("RemedyType").
		Preload("UserDentist").Find(&treatments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatments})
}

func GetTreatments(c *gin.Context) {
	var treatments entity.Treatment
	id := c.Param("id")
	if err := entity.DB().Preload("Screening").
		Preload("Screening.Patient").
		Raw("SELECT * FROM treatments WHERE id = ? ", id).Find(&treatments).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}
	c.JSON(http.StatusOK, gin.H{"data": treatments})

}