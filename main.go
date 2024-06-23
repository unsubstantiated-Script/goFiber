package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"goFiber/database"
	"goFiber/lead"
)

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	err := app.Listen("3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(DBConn *gorm.DB) {
		err := DBConn.Close()
		if err != nil {

		}
	}(database.DBConn)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Migrated leads")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}
