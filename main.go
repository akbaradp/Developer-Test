package main

import (
	"log"
	"net/http"
	"user_management/handler"
	"user_management/repository"
	"user_management/service"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Setup database connection
	db, err := gorm.Open(postgres.Open("user=postgres password=1234 dbname=postgres sslmode=disable search_path=golang"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the database
	err = repository.Migrate(db)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repositories
	nationalityRepo := repository.NewNationalityRepository(db)
	customersRepo := repository.NewCustomersRepository(db)
	familyRepo := repository.NewFamilyRepository(db)

	// Initialize services
	nationalityService := service.NewNationalityService(nationalityRepo)
	customersService := service.NewCustomersService(customersRepo)
	familyService := service.NewFamilyService(familyRepo)

	// Initialize handlers
	nationalityHandler := handler.NewNationalityHandler(nationalityService)
	customerHandler := handler.NewCustomerHandler(customersService)
	familyHandler := handler.NewFamilyHandler(familyService)

	// Setup routes
	r := mux.NewRouter()

	r.HandleFunc("/nationalities", nationalityHandler.GetAllNationalities).Methods("GET")
	r.HandleFunc("/customers", customerHandler.GetAllCustomers).Methods("GET")
	r.HandleFunc("/families", familyHandler.GetAllFamilies).Methods("GET")

	// Start server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
