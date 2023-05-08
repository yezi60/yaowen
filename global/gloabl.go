package global

import "gorm.io/gorm"

var (
	DB *gorm.DB
)

const (
	ADDRCHANNUM = 10
	HEALCHANNUM = 10
	WORKERNUM = 5
)