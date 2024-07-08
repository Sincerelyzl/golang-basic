package main

import (
	"fmt"
	"gorm/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {

	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.Movie{})
	fmt.Println("Database migration completed!")

	//Create a new movie
	// newMovie := &models.Movie{
	// 	Name:        "The BankNa",
	// 	Director:    "BankNa",
	// 	Description: "A movie about a BankNa",
	// 	Score:       1,
	// }
	// models.CreateMovie(db, newMovie)

	// Get a movie by ID
	// playMovie := models.GetMovies(db, 1)

	// playMovie.Name = "The Phancer 2"
	// playMovie.Director = "PapsinV2"
	// playMovie.Description = "A movie about a phancer V2"
	// playMovie.Score = 9

	// Update a movie
	// models.UpdateMovie(db, playMovie)

	// Delete a movie
	models.DeleteMovie(db, 2)

}
