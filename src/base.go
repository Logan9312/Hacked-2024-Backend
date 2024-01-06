package src

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

var DB *sqlx.DB

type Payment struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Price       int64  `db:"price" json:"price"`
}

func FetchPayments(c echo.Context) error {

	if DB == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection is not established"})
	}

	var payments []Payment
	err := DB.Select(&payments, "SELECT * FROM payments")

	if err != nil {
		// Log the detailed error for internal use
		log.Printf("Error fetching products: %v", err)

		// Return a generic error message to the client
		return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Error fetching products"})
	}

	return c.JSON(http.StatusOK, payments)

}
