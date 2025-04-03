package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/apex/gateway/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Initialize the Echo application
func setupRouter() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", Hello)
	e.GET("/get-student", GetStudent)
	e.GET("/get-grade", GetGrade)
	e.GET("/get-teacher", GetTeacher)

	return e
}

// Handler
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetStudent(c echo.Context) error {
	return c.String(http.StatusOK, "Test API Get Student.")
}

func GetGrade(c echo.Context) error {
	return c.String(http.StatusOK, "Test API Get Grade.")
}

func GetTeacher(c echo.Context) error {
	return c.String(http.StatusOK, "Test API Get Teacher.")
}

func main() {
	// Set up the Echo router
	e := setupRouter()

	// Check if we're running locally
	isLocal := false
	if localEnv, exists := os.LookupEnv("IS_LOCAL"); exists {
		isLocal, _ = strconv.ParseBool(localEnv)
	}

	if isLocal {
		// Start the Echo server locally
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		e.Logger.Printf("Starting local server on port %s", port)
		e.Logger.Fatal(e.Start(":" + port))
	} else {
		// Start with Apex Gateway
		e.Logger.Print("Starting Lambda handler with Apex Gateway")

		// Use Apex Gateway to handle Lambda events
		gateway.ListenAndServe("", e)
	}
}
