package model

type Data struct {
	Address string    `gorm:"primarykey;type:varchar(200);not null"`
	Health int32      `gorm:"type:int;not null"`
}
