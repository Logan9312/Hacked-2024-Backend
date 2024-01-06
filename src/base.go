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

type Task struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	DueDate     string `db:"due_date" json:"due_date"`
	AssignedTo  string `db:"assigned_to" json:"assigned_to"`
	Completed   bool   `db:"completed" json:"completed"`
}

type List struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

type Message struct {
	ID        int64  `db:"id" json:"id"`
	Content   string `db:"content" json:"content"`
	Author    string `db:"author" json:"author"`
	Timestamp string `db:"timestamp" json:"timestamp"`
}

type User struct {
	ID    int64  `db:"id" json:"id"`
	Name  string `db:"username" json:"username"`
	Email string `db:"email" json:"email"`
}

func FetchPayments(c echo.Context) error {

	if DB == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection is not established"})
	}

	var payments []Payment
	err := DB.Select(&payments, "SELECT * FROM payments")

	if err != nil {
		// Log the detailed error for internal use
		log.Printf("Error fetching payments: %v", err)

		// Return a generic error message to the client
		return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Error fetching payments"})
	}

	return c.JSON(http.StatusOK, payments)
}

func SavePayment(c echo.Context) error {
	var newPayment Payment
	if err := c.Bind(&newPayment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	query := "INSERT INTO payments (name, description, price) VALUES (:name, :description, :price)"
	result, err := DB.NamedExec(query, newPayment)
	if err != nil {
		// Step 4: Handle any errors
		log.Printf("Error saving payment: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error saving payment"})
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting new payment ID: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error getting new payment ID"})
	}
	newPayment.ID = id

	return c.JSON(http.StatusCreated, newPayment)
}

func FetchTasks(c echo.Context) error {
	if DB == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection is not established"})
	}

	var tasks []Task
	err := DB.Select(&tasks, "SELECT * FROM tasks")
	if err != nil {
		log.Printf("Error fetching tasks: %v", err)
		return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Error fetching tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

// SaveTask saves a new task to the database.
func SaveTask(c echo.Context) error {
	var newTask Task
	if err := c.Bind(&newTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	query := "INSERT INTO tasks (name, description, due_date, assigned_to, completed) VALUES (:name, :description, :due_date, :assigned_to, :completed)"
	result, err := DB.NamedExec(query, newTask)
	if err != nil {
		log.Printf("Error saving task: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error saving task"})
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting new task ID: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error getting new task ID"})
	}
	newTask.ID = id

	return c.JSON(http.StatusCreated, newTask)
}

// FetchLists retrieves all lists from the database.
func FetchLists(c echo.Context) error {
	if DB == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection is not established"})
	}

	var lists []List
	err := DB.Select(&lists, "SELECT * FROM lists")
	if err != nil {
		log.Printf("Error fetching lists: %v", err)
		return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Error fetching lists"})
	}

	return c.JSON(http.StatusOK, lists)
}

// SaveList saves a new list to the database.
func SaveList(c echo.Context) error {
	var newList List
	if err := c.Bind(&newList); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	query := "INSERT INTO lists (name, description) VALUES (:name, :description)"
	result, err := DB.NamedExec(query, newList)
	if err != nil {
		log.Printf("Error saving list: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error saving list"})
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting new list ID: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error getting new list ID"})
	}
	newList.ID = id

	return c.JSON(http.StatusCreated, newList)
}

// FetchMessages retrieves all messages from the database.
func FetchMessages(c echo.Context) error {
	if DB == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection is not established"})
	}

	var messages []Message
	err := DB.Select(&messages, "SELECT * FROM messages")
	if err != nil {
		log.Printf("Error fetching messages: %v", err)
		return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Error fetching messages"})
	}

	return c.JSON(http.StatusOK, messages)
}

// SaveMessage saves a new message to the database.
func SaveMessage(c echo.Context) error {
	var newMessage Message
	if err := c.Bind(&newMessage); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	query := "INSERT INTO messages (content, author, timestamp) VALUES (:content, :author, :timestamp)"
	result, err := DB.NamedExec(query, newMessage)
	if err != nil {
		log.Printf("Error saving message: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error saving message"})
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting new message ID: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error getting new message ID"})
	}
	newMessage.ID = id

	return c.JSON(http.StatusCreated, newMessage)
}

// FetchUsers retrieves all users from the database.
func FetchUsers(c echo.Context) error {
	if DB == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection is not established"})
	}

	var users []User
	err := DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Error fetching users"})
	}

	return c.JSON(http.StatusOK, users)
}

// SaveUser saves a new user to the database.
func SaveUser(c echo.Context) error {
	var newUser User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	query := "INSERT INTO users (username, email) VALUES (:username, :email)"
	result, err := DB.NamedExec(query, newUser)
	if err != nil {
		log.Printf("Error saving user: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error saving user"})
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting new user ID: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error getting new user ID"})
	}
	newUser.ID = id

	return c.JSON(http.StatusCreated, newUser)
}
