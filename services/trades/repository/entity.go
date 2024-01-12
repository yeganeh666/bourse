package repository

import "time"

type Instrument struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

type Trade struct {
	ID           int       `gorm:"primaryKey"`
	InstrumentID int       `gorm:"index"`
	DateEn       time.Time `gorm:"type:timestamptz"`
	Open         float64   `gorm:"type:decimal"`
	High         float64   `gorm:"type:decimal"`
	Low          float64   `gorm:"type:decimal"`
	Close        float64   `gorm:"type:decimal"`
}
