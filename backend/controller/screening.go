package controller

import (
	"github.com/cottoncandyblue/sa-65-project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /screening
func CreateScreening(c *gin.Context) {

	var screening_record entity.Screening
	var patient entity.Patient
	var medical_product entity.MedicalProduct
	var userdentistass entity.User

	//10:ผลลัพธ์ที่ได้จากขั้นตอนที่ x จะถูก bind เข่้าตัวแปร scr
	if err := c.ShouldBindJSON(&screening_record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//31:ค้นหา User ด้วย id
	if tx := entity.DB().Where("id = ?", screening_record.UserDentistassID).First(&userdentistass); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dentist not found"})
		return
	}

	
	entity.DB().Joins("Role").Find(&userdentistass)

	if userdentistass.Role.Name != "Dentist" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "only for dentsit"})
		return
	}
	//7:ค้นหา patient ด้วย p_id
	if tx := entity.DB().Where("id = ?", screening_record.PatientID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Patient not found"})
		return
	}
	//11:ค้นหา medical_product ด้วย m_id
	if tx := entity.DB().Where("id = ?", screening_record.MedicalProductID).First(&medical_product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Medical Product not found"})
		return
	}
	//12:สร้าง Screening_records(p_id, m_id, u_id, illnesses, detail)
	scr := entity.Screening{
		//โยงความสัมพันธ์กับ Entity Patient
		//โยงความสัมพันธ์กับ Entity Medical_product
		//โยงความสัมพันธ์กับ Entity User
		Patient:        patient,
		MedicalProduct: medical_product,
		UserDentistass: userdentistass,
		Illnesses:      screening_record.Illnesses,
		//Detail:         screening_record.Detail,
		Queue:          screening_record.Queue,
	}
	//13:บันทึก()
	if err := entity.DB().Create(&scr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": scr})
}

// GET /screening
func GetScreening(c *gin.Context) {
	var screening_record entity.Screening
	id := c.Param("id")
	if err := entity.DB().Preload("Patient").Preload("MedicalProduct").Preload("User").Raw("SELECT * FROM screenings WHERE id = ?", id).Find(&screening_record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": screening_record})
}

// GET /screenings
func ListScreening(c *gin.Context) {
	var screening_records []entity.Screening
	if err := entity.DB().Preload("Patient").Preload("MedicalProduct").Raw("SELECT * FROM screenings").Find(&screening_records).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": screening_records})
}
