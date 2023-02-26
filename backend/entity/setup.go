package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Role{}, &User{},
		&Job{}, &Insurance{}, &Patient{}, &Sex{},
		&Screening{},
		&Treatment{}, &RemedyType{},
		&Appoint{},
		&MedicalProduct{}, &MedRecord{},
		&Payment{},
	)

	db = database

	// ตำแหน่งงาน --------------------------------------------------------------
	role1 := Role{
		Name: "Dentist",
	}
	db.Model(&Role{}).Create(&role1)

	role2 := Role{
		Name: "Dental assistant",
	}
	db.Model(&Role{}).Create(&role2)

	role3 := Role{
		Name: "Nurse",
	}
	db.Model(&Role{}).Create(&role3)

	role4 := Role{
		Name: "Pharmacist",
	}
	db.Model(&Role{}).Create(&role4)

	role5 := Role{
		Name: "Financial officer",
	}
	db.Model(&Role{}).Create(&role5)

	// รวมสมาชิกทุกตำแหน่ง >> entity User -------------------------------------------------------
	password1, err := bcrypt.GenerateFromPassword([]byte("1234"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("5678"), 14)
	dentist1 := User{
		Name:     "กอเอ๋ย กอไก่",
		Username: "khana",
		Password: string(password2),
		Role:     role1,
	}
	db.Model(&User{}).Create(&dentist1)

	dentist2 := User{
		Name:     "ขอไข่ ในเล้า",
		Username: "name",
		Password: string(password1),
		Role:     role1,
	}
	db.Model(&User{}).Create(&dentist2)

	dentistass1 := User{
		Name:     "คอควาย เข้านา",
		Username: "a5678",
		Password: string(password1),
		Role:     role2,
	}
	db.Model(&User{}).Create(&dentistass1)

	dentistass2 := User{
		Name:     "งองู ใจกล้า",
		Username: "d5678",
		Password: string(password2),
		Role:     role2,
	}
	db.Model(&User{}).Create(&dentistass2)

	nurse1 := User{
		Name:     "จอจาน ใช้ดี",
		Username: "w1234",
		Password: string(password1),
		Role:     role3,
	}
	db.Model(&User{}).Create(&nurse1)

	nurse2 := User{
		Name:     "ฉอฉิ่ง ตีดัง",
		Username: "pcrc",
		Password: string(password2),
		Role:     role3,
	}
	db.Model(&User{}).Create(&nurse2)

	pharmacist1 := User{
		Name:     "ชอช้าง วิ่งหนี",
		Username: "l1234",
		Password: string(password1),
		Role:     role4,
	}
	db.Model(&User{}).Create(&pharmacist1)

	pharmacist2 := User{
		Name:     "ซอโซ่ ล่ามดี",
		Username: "q1234",
		Password: string(password2),
		Role:     role4,
	}
	db.Model(&User{}).Create(&pharmacist2)

	financial1 := User{
		Name:     "ญอหญิง โสภา",
		Username: "padthai",
		Password: string(password1),
		Role:     role5,
	}
	db.Model(&User{}).Create(&financial1)

	financial2 := User{
		Name:     "ฐอฐาน เข้ามารอง",
		Username: "s1234",
		Password: string(password2),
		Role:     role5,
	}
	db.Model(&User{}).Create(&financial2)

	// เพศ --------------------------------------------------------------------------------------
	sex1 := Sex{
		Name: "ชาย",
	}
	db.Model(&Sex{}).Create(&sex1)

	sex2 := Sex{
		Name: "หญิง",
	}
	db.Model(&Sex{}).Create(&sex2)

	// อาชีพ --------------------------------------------------------------------------------------------
	job1 := Job{
		Name: "ราชการ",
	}
	db.Model(&Job{}).Create(&job1)

	job2 := Job{
		Name: "รัฐวิสหกิจ",
	}
	db.Model(&Job{}).Create(&job2)

	job3 := Job{
		Name: "นักศึกษา",
	}
	db.Model(&Job{}).Create(&job3)

	// สิทธิการรักษา ------------------------------------------------------------------------------------------------
	insurance1 := Insurance{
		Name:   "สิทธิสวัสดิการข้าราชการ",
		Detail: "ข้าราชการและบคุคลในครอบครัวสามารถใช้สิทธิ์เบิกจ่ายตรง โดยใช้บัตรประชาชนในการเข้ารับบริการรักษาพยาบาลประเภทผู้ป่วยนอกทุกครั้ง ณ จุดชำระเงินโดยหากไม่ได้นำบัตรประชาชนมาแสดง หรือเอกสารที่กรมบัญชีกลางกำหนด ผู้รับบริการจะต้องสำรองจ่ายเงินค่ารักษาพยาบาลไปก่อน แล้วนำใบเสร็จรับเงินไปเบิกคืนกับส่วนราชการต้นสังกัด",
	}
	db.Model(&Insurance{}).Create(&insurance1)

	insurance2 := Insurance{
		Name:   "สิทธิประกันสังคม",
		Detail: "สามารถใช้สิทธิ์ได้เฉพาะกรณีที่มีใบส่งตัวมาจากโรงพยาบาลต้นสังกัด และชำระเงินสดเท่านั้น ยกเว้น กรณีมีใบส่งตัวยืนยันการให้วางบิลโรงพยาบาลต้นสังกัดได้ ",
	}
	db.Model(&Insurance{}).Create(&insurance2)

	insurance3 := Insurance{
		Name:   "สิทธิหลักประกันสุขภาพ 30 บาท",
		Detail: "คุ้มครองบุคคลที่เป็นคนไทยมีเลขประจำตัวประชาชน 13 หลักที่ไม่ได้รับสิทธิสวัสดิการข้าราชการ หรือ สิทธิประกันสังคม หรือสิทธิสวัสดิการรัฐวิสาหกิจ หรือสิทธิอื่น ๆ จากรัฐ",
	}
	db.Model(&Insurance{}).Create(&insurance3)

	// ประเภทการรักษา ------------------------------------------------------------------------------------------------
	remedy1 := RemedyType{
		Name: "อุดฟัน",
	}
	db.Model(&RemedyType{}).Create(&remedy1)

	remedy2 := RemedyType{
		Name: "ขูดหินปูน",
	}
	db.Model(&RemedyType{}).Create(&remedy2)

	remedy3 := RemedyType{
		Name: "เอ็กซ์เรย์",
	}
	db.Model(&RemedyType{}).Create(&remedy3)

	// เวชระเบียน ------------------------------------------------------------------------------------------------------------
	patient1 := Patient{
		Firstname: "ข้าวหอม",
		Lastname:  "สดใส",
		Age:       20,
		
		Sex:       sex1,
		Job:       job3,
		Insurance: insurance3,
		UserNurse: nurse1,
	}
	db.Model(&Patient{}).Create(&patient1)

	patient2 := Patient{
		Firstname: "สมหญิง",
		Lastname:  "ซิ่งรถไถ",
		Age:       26,
	
		Sex:       sex2,
		Job:       job1,
		Insurance: insurance1,
		UserNurse: nurse1,
	}
	db.Model(&Patient{}).Create(&patient2)

	patient3 := Patient{
		Firstname: "สมชาย",
		Lastname:  "มาอุดฟัน",
		Age:       57,
		
		Sex:       sex2,
		Job:       job1,
		Insurance: insurance1,
		UserNurse: nurse2,
	}
	db.Model(&Patient{}).Create(&patient3)

	// ยาและเวชภัณฑ์ -----------------------------------------------------------------------------------------------
	MedicalProduct1 := MedicalProduct{
		Name: "Paracetamol(กระปุก)",
	}
	db.Model(&MedicalProduct{}).Create(&MedicalProduct1)

	MedicalProduct2 := MedicalProduct{
		Name: "Paracetamol(เม็ด)",
	}
	db.Model(&MedicalProduct{}).Create(&MedicalProduct2)

	MedicalProduct3 := MedicalProduct{
		Name: "ไหมขัดฟัน",
	}
	db.Model(&MedicalProduct{}).Create(&MedicalProduct3)

	// คัดกรองผู้ป่วย ------------------------------------------------------------------------------------------------
	screening1 := Screening{
		Illnesses:      "ปวดฟัน",
		//Detail:         "ปวดฟันมานาน 1 ชั่วโมง",
		Queue:          "A10",
		Patient:        patient1,
		UserDentistass: dentistass1,
		MedicalProduct: MedicalProduct2,
	}
	db.Model(&Screening{}).Create(&screening1)

	screening2 := Screening{
		Illnesses:      "เหงือกอักเสบ",
		//Detail:         "มีอาการเหงือกบวม",
		Queue:          "A11",
		Patient:        patient2,
		UserDentistass: dentistass1,
		MedicalProduct: MedicalProduct1,
	}
	db.Model(&Screening{}).Create(&screening2)

	screening3 := Screening{
		Illnesses:      "ปวดฟัน",
	//	Detail:         "ปวดฟันมานาน 2 ชั่วโมง",
		Queue:          "A12",
		Patient:        patient3,
		UserDentistass: dentistass1,
		MedicalProduct: MedicalProduct2,
	}
	db.Model(&Screening{}).Create(&screening3)

	// ใบวินิฉัย --------------------------------------------------------------------------------------------------------------------
	treatment1 := Treatment{
		PrescriptionRaw:  "A12",
		PrescriptionNote: "",
		ToothNumber:      "21",
		Date:             time.Now(),
		Screening:        screening1,
		UserDentist:      dentist1,
		RemedyType:       remedy1,
	}
	db.Model(&Treatment{}).Create(&treatment1)

	treatment2 := Treatment{
		PrescriptionRaw:  "A12",
		PrescriptionNote: "",
		ToothNumber:      "21",
		Date:             time.Now(),
		Screening:        screening2,
		UserDentist:      dentist1,
		RemedyType:       remedy2,
	}
	db.Model(&Treatment{}).Create(&treatment2)

	treatment3 := Treatment{
		PrescriptionRaw:  "A12",
		PrescriptionNote: "",
		ToothNumber:      "21",
		Date:             time.Now(),
		Screening:        screening3,
		UserDentist:      dentist2,
		RemedyType:       remedy3,
	}
	db.Model(&Treatment{}).Create(&treatment3)

	// การนัดหมาย ---------------------------------------------------------------------------------------------------------------
	
	// รายการบันทึกการจ่ายยา ------------------------------------------------------------------------------------------------------------
	MedRecord1 := MedRecord{
		Amount:         2,
		Treatment:      treatment1,
		UserPharmacist: pharmacist1,
		MedicalProduct: MedicalProduct2,
	}
	db.Model(&MedRecord{}).Create(&MedRecord1)

	MedRecord2 := MedRecord{
		Amount:         2,
		Treatment:      treatment2,
		UserPharmacist: pharmacist1,
		MedicalProduct: MedicalProduct1,
	}
	db.Model(&MedRecord{}).Create(&MedRecord2)

	MedRecord3 := MedRecord{
		Amount:         3,
		Treatment:      treatment3,
		UserPharmacist: pharmacist2,
		MedicalProduct: MedicalProduct3,
	}
	db.Model(&MedRecord{}).Create(&MedRecord3)

	// การชำระเงิน ------------------------------------------------------------------------------------------------
	Payment1 := Payment{
		Price:         2000.00,
		Paytime:       time.Now(),
		Note:          "",
		Patient:       patient1,
		UserFinancial: financial1,
		RemedyType:    remedy1,
	}
	db.Model(&Payment{}).Create(&Payment1)

	Payment2 := Payment{
		Price:         200.00,
		Paytime:       time.Now(),
		Note:          "",
		Patient:       patient2,
		UserFinancial: financial1,
		RemedyType:    remedy2,
	}
	db.Model(&Payment{}).Create(&Payment2)

	Payment3 := Payment{
		Price:         500.00,
		Paytime:       time.Now(),
		Note:          "",
		Patient:       patient3,
		UserFinancial: financial1,
		RemedyType:    remedy3,
	}
	db.Model(&Payment{}).Create(&Payment3)

}
