package entity

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	Date        time.Time
	Description string
	Amount      float32
	Category    string
	Account     string
}
