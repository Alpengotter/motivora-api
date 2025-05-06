package main

import (
	"fmt"
	"log"
	"motivora-backend/internal/handlers"
	"motivora-backend/internal/middleware"
	"net/http"
)

func main() {
	fmt.Println("Server started at :8080")

	// Обработчик логина
	http.Handle("/api/v1/login", middleware.ApplyMiddlewares(http.HandlerFunc(handlers.LoginHandler)))

	http.Handle("/api/v1/employers/", http.StripPrefix("/api/v1/employers", middleware.ApplyMiddlewares(handlers.EmployerRouter())))
	http.Handle("/api/v1/companies/", http.StripPrefix("/api/v1/companies", middleware.ApplyMiddlewares(handlers.CompanyRouter())))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
