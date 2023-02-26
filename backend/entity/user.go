package entity

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name  string
	Users []User `gorm:"foreignKey:RoleID"`
}

type User struct {
	gorm.Model
	Name     string
	Username string `gorm:"uniqueIndex"`
	Password string

	Patients   []Patient   `gorm:"foreignKey:UserNurseID"`
	Screenings []Screening `gorm:"foreignKey:UserDentistassID"`
	Treatments []Treatment `gorm:"foreignKey:UserDentistID"`
	Appoints   []Appoint   `gorm:"foreignKey:UserDentistID"`
	MedRecords []MedRecord `gorm:"foreignKey:UserPharmacistID"`
	Payments   []Payment   `gorm:"foreignKey:UserFinancialID"`

	RoleID *uint
	Role   Role
}

//ระบบย่อย ระบบบันทึกเวชระเบียน
type Job struct {
	gorm.Model
	Name string

	Patients []Patient `gorm:"foreignKey:JobID"`
}

type Insurance struct {
	gorm.Model
	Name   string
	Detail string

	Patients []Patient `gorm:"foreignKey:InsuranceID"`
}

type Sex struct {
	gorm.Model
	Name string

	Patients []Patient `gorm:"foreignKey:SexID"`
}

type Patient struct {
	gorm.Model
	Firstname string
	Lastname  string
	Age       int
	

	UserNurseID *uint
	UserNurse   User

	JobID *uint
	Job   Job

	InsuranceID *uint
	Insurance   Insurance

	SexID *uint
	Sex   Sex

	Screenings []Screening `gorm:"foreignKey:PatientID"`
	Appoints   []Appoint   `gorm:"foreignKey:PatientID"`
	Payments   []Payment   `gorm:"foreignKey:PatientID"`
}

//ระบบย่อย ระบบคัดกรองข้อมูลพื้นฐานผู้ป่วย

type Screening struct {
	gorm.Model
	Illnesses string
	//Detail    string
	Queue     string

	PatientID *uint
	Patient   Patient

	UserDentistassID *uint
	UserDentistass   User

	MedicalProductID *uint
	MedicalProduct   MedicalProduct
}

//ระบบย่อย ระบบบันทักการรักษาทางทันตกรรม
type Treatment struct {
	gorm.Model

	PrescriptionRaw  string
	PrescriptionNote string
	ToothNumber      string
	Date             time.Time

	ScreeningID *uint
	Screening   Screening

	UserDentistID *uint
	UserDentist   User

	RemedyTypeID *uint
	RemedyType   RemedyType

	MedRecords []MedRecord `gorm:"foreignKey:TreatmentID"`
}

//ระบบย่อย ระบบบันทึกการนัดหมาย
type RemedyType struct {
	gorm.Model
	Name      string
	Appoints  []Appoint   `gorm:"foreignKey:RemedyTypeID"`
	Treatment []Treatment `gorm:"foreignKey:RemedyTypeID"`
	Payments  []Payment   `gorm:"foreignKey:RemedyTypeID"`
}

type Appoint struct {
	gorm.Model
	AppointTime time.Time
	Todo        string

	UserDentistID *uint
	UserDentist   User

	PatientID *uint
	Patient   Patient

	RemedyTypeID *uint
	RemedyType   RemedyType
}

//ระบบย่อย ระบบบันทึกการจ่ายยาและเวชภัณฑ์
type MedicalProduct struct {
	gorm.Model
	Name       string
	Screenings []Screening `gorm:"foreignKey:MedicalProductID"`
	MedRecords []MedRecord `gorm:"foreignKey:MedicalProductID"`
}

type MedRecord struct {
	gorm.Model
	Amount uint

	TreatmentID *uint
	Treatment   Treatment

	UserPharmacistID *uint
	UserPharmacist   User

	MedicalProductID *uint
	MedicalProduct   MedicalProduct
}

//ระบบย่อย ระบบบันทึกการชำระเงิน
type Payment struct {
	gorm.Model

	Price   float32
	Paytime time.Time
	Note    string

	PatientID *uint
	Patient   Patient

	UserFinancialID *uint
	UserFinancial   User

	RemedyTypeID *uint
	RemedyType   RemedyType
}
