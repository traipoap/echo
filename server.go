package main

import (
	"log"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CustomerHandler struct {
	DB *gorm.DB // Field DB to Database
}

func (h *CustomerHandler) Initialize() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Customer{})
	h.DB = db
}

type Customer struct {
	Id        uint   `gorm:"primary_key" json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

func (h *CustomerHandler) GetAllCustomer(c echo.Context) error {
	customers := []Customer{}

	h.DB.Find(&customers)

	return c.JSON(http.StatusOK, customers)
}

func (h *CustomerHandler) GetCustomer(c echo.Context) error {
	id := c.Param("id")
	customer := Customer{}

	if err := h.DB.Find(&customer, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) SaveCustomer(c echo.Context) error {
	customer := Customer{}

	if err := c.Bind(&customer); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&customer).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
	id := c.Param("id")
	customer := Customer{}

	if err := h.DB.Find(&customer, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := c.Bind(&customer); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&customer).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) DeleteCustomer(c echo.Context) error {
	id := c.Param("id")
	customer := Customer{}

	if err := h.DB.Find(&customer, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := h.DB.Delete(&customer).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func main() {

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, remote_ip=${remote_ip}, method=${method}, host=${host}, uri=${uri}, status=${status}, error=${error}, latency_human=${latency_human}\n",
	}))

	h := CustomerHandler{}
	h.Initialize()

	customerGroup := e.Group("/customers")
	customerGroup.POST("", h.SaveCustomer)     // C
	customerGroup.GET("", h.GetAllCustomer)    // R
	customerGroup.GET("/:id", h.GetCustomer)   // R ID
	customerGroup.PUT("", h.UpdateCustomer)    // U
	customerGroup.DELETE("", h.DeleteCustomer) // D

	e.Logger.Fatal(e.Start(":8080"))
}
