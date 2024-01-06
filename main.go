package Main

import (
	"fmt"
	"log"
	"os"

	"github.com/Logan9312/Hacked-2024-Backend/src"
	"github.com/Logan9312/Hacked-2024-Backend/routers"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	stripe.Key = os.Getenv("STRIPE_TOKEN")
	connStr := os.Getenv("DATABASE_URL")

	src.DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer src.DB.Close()

	rows, err := src.DB.Query("select version()")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var version string
	for rows.Next() {
		err := rows.Scan(&version)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("version=%s\n", version)

	//go commands.SetRoles(mainSession)
	fmt.Println("Backend is running!")

	routers.HealthCheck()

}
