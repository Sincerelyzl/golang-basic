package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Name        string
	Director    string
	Description string
	Score       uint
}

func CreateMovie(db *gorm.DB, movie *Movie) {
	result := db.Create(movie)

	if result.Error != nil {
		log.Fatalf("Unable to create movie: %v", result.Error)
	}
	fmt.Println("Movie created successfully!")
}

func GetMovies(db *gorm.DB, id uint) *Movie {
	var movie Movie
	result := db.First(&movie, id)
	if result.Error != nil {
		log.Fatalf("Unable to play movie: %v", result.Error)
	}
	return &movie
}

func UpdateMovie(db *gorm.DB, movie *Movie) {
	result := db.Save(&movie)
	if result.Error != nil {
		log.Fatalf("Unable to remake movie: %v", result.Error)
	}
	fmt.Println("Movie remake successfully!")
}

func DeleteMovie(db *gorm.DB, id uint) {
	var movie Movie
	result := db.Unscoped().Delete(&movie, id)

	if result.Error != nil {
		log.Fatalf("Unable to Cancel movie: %v", result.Error)
	}
	fmt.Println("Movie Cancelled successfully!")
}
