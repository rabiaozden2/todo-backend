package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rabiaozden/todo-backend/internal/models"
)

func Connect() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL bos, .env dosyasini kontrol et")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("veritabanina baglanilamadi: %v", err)
	}

	log.Println("veritabani baglantisi kuruldu")

	// Otomatik tablo olusturma (Migrations)
	err = database.AutoMigrate(
		&models.User{},
		&models.Task{},
	)
	if err != nil {
		log.Printf("Tablo olusturma hatasi: %v", err)
	}

	return database
}