package models

import "gorm.io/gorm"

// Gunakan schema "golang"
func (Nationality) TableName(tx *gorm.DB) string {
	return "golang.nationalit"
}

func (Customer) TableName(tx *gorm.DB) string {
	return "golang.customer"
}

func (FamilyList) TableName(tx *gorm.DB) string {
	return "golang.family_list"
}

type Nationality struct {
	NationalityID   int    `gorm:"primaryKey"`
	NationalityName string `gorm:"type:varchar(50);not null"`
	NationalityCode string `gorm:"type:char(2);not null"`
}

type Customer struct {
	CstID         int          `gorm:"primaryKey"`
	NationalityID int          `gorm:"not null"`
	CstName       string       `gorm:"type:varchar(50);not null"`
	CstDOB        string       `gorm:"type:date;not null"`
	CstPhonenum   string       `gorm:"type:varchar(20);not null"`
	CstEmail      string       `gorm:"type:varchar(50);not null"`
	Nationality   Nationality  `gorm:"foreignKey:NationalityID"`
	FamilyList    []FamilyList `gorm:"foreignKey:CstID"`
}

type FamilyList struct {
	FlID       int      `gorm:"primaryKey"`
	CstID      int      `gorm:"not null"`
	FlRelation string   `gorm:"type:varchar(50);not null"`
	FlName     string   `gorm:"type:varchar(50);not null"`
	FlDOB      string   `gorm:"type:varchar(50);not null"`
	Customer   Customer `gorm:"foreignKey:CstID"`
}
