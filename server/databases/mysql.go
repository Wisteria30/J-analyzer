package databases

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (db *gorm.DB, err error) {
	err = godotenv.Load()

	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+
			"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+
			os.Getenv("DB_DATABASE")+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		logrus.Fatal(err)
	}

	return db, err
}