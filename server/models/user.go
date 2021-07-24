package models

import "time"

type User struct {
	ID uint `gorm:"primary_key"`
	UID string `json:"-"`
	AccessToken string `json:"access_token"`
	Email string `json:"-"`
	GyazoUID string `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}