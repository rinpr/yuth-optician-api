package models

type CustomerV1 struct {
	CustomerID   int             `gorm:"type:int;primary_key"`
	PersonalData PersonalDataV1  `gorm:"type:varchar(255)"`
	Bill         []Bill          `gorm:"type:varchar(255)"`
}

type PersonalDataV1 struct {
	FirstName string `gorm:"type:varchar(255)"`
	LastName  string `gorm:"type:varchar(255)"`
	Address   string `gorm:"type:varchar(255)"`
	Phone     string `gorm:"type:varchar(255)"`
	BirthDate string `gorm:"type:varchar(255)"`
	Gender    string `gorm:"type:varchar(255)"`
	Picture   string `gorm:"type:varchar(255)"`
}
