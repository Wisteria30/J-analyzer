package main

import (
	"github.com/Wisteria30/J-analyzer/databases"
	"github.com/Wisteria30/J-analyzer/models"
	"github.com/sirupsen/logrus"
)

func main() {
	d, err := databases.Connect()
	if err != nil {
		logrus.Fatal(err)
	}
	db, _ := d.DB()
	defer db.Close()

	d.Debug().AutoMigrate(&models.User{})
}