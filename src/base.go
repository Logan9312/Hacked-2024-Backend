package src

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

var DB *sqlx.DB

type HouseHold struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Payment struct {
	ID          int64  `db:"id" json:"id"`
	HouseHoldID int64  `db:"household" json:"household"`
	Payee       string `db:"payee" json:"payee"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Price       int64  `db:"price" json:"price"`
}

type Task struct {
	ID          int64  `db:"id" json:"id"`
	HouseHoldID int64  `db:"household" json:"household"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	DueDate     string `db:"due_date" json:"due_date"`
	AssignedTo  string `db:"assigned_to" json:"assigned_to"`
	Completed   bool   `db:"completed" json:"completed"`
}

type List struct {
	ID          int64  `db:"id" json:"id"`
	HouseHold   int64  `db:"household" json:"household"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

type Message struct {
	ID        int64  `db:"id" json:"id"`
	HouseHold int64  `db:"household" json:"household"`
	Content   string `db:"content" json:"content"`
	AuthorID  int64  `db:"author_id" json:"author_id"`
	Timestamp string `db:"timestamp" json:"timestamp"`
}

type AppUser struct {
	ID        int64  `db:"id" json:"id"`
	HouseHold int64  `db:"household" json:"household"`
	Username  string `db:"username" json:"username"`
	Email     string `db:"email" json:"email"`
}

type Items struct {
	ID        int64  `db:"id" json:"id"`
	HouseHold int64  `db:"household" json:"household"`
	ListID    int64  `db:"list_id" json:"list_id"`
	Name      string `db:"name" json:"name"`
}

func FetchPayments(c echo.Context) error {
	if DB == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection is not established"})
	}

	var payments []Payment
	err := DB.Select(&payments, "SELECT * FROM payment")
	if err != nil {
		log.Printf("Error fetching payments: %v", err)
		return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Error fetching payments"})
	}

	for k, v := range payments {
		err := DB.Get(&v.Payee, "SELECT name FROM appuser WHERE id = $1", v.Payee)
		if err != nil {
			log.Printf("Error fetching user: %v", err)
			payments[k].Payee = "Unknown"
		}

		//payments[k].Payee = v.Name

	}

	return c.JSON(http.StatusOK, payments)
}

func SavePayment(c echo.Context) error {
	var newPayment Payment
	if err := c.Bind(&newPayment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	query := "INSERT INTO payment (household_id, payee, name, description, price) VALUES (:household_id, :payee, :name, :description, :price)"
	result, err := DB.NamedExec(query, newPayment)
	if err != nil {
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
	err := DB.Select(&tasks, "SELECT * FROM task")
	if err != nil {
		log.Printf("Error fetching task: %v", err)
		return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Error fetching tasks"})
	}

	for k, v := range tasks {
		err := DB.Get(&v.AssignedTo, "SELECT username FROM appuser WHERE id = $1", v.AssignedTo)
		if err != nil {
			log.Printf("Error fetching user: %v", err)
			tasks[k].AssignedTo = "Unknown"
		}

		//tasks[k].AssignedTo = v.Name

	}

	return c.JSON(http.StatusOK, tasks)
}

// SaveTask saves a new task to the database.
func SaveTask(c echo.Context) error {
	var newTask Task
	if err := c.Bind(&newTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	query := "INSERT INTO task (household_id, name, description, due_date, assigned_to_id, completed) VALUES (:household_id, :name, :description, :due_date, :assigned_to_id, :completed)"
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
	err := DB.Select(&lists, "SELECT * FROM list")
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

	query := "INSERT INTO list (household_id, name, description) VALUES (:household_id, :name, :description)"
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
	err := DB.Select(&messages, "SELECT * FROM message")
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

	query := "INSERT INTO message (household_id, content, author_id, timestamp) VALUES (:household_id, :content, :author_id, :timestamp)"
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

	var users []AppUser
	err := DB.Select(&users, "SELECT * FROM appuser")
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Error fetching users"})
	}

	return c.JSON(http.StatusOK, users)
}

// SaveUser saves a new user to the database.
func SaveUser(c echo.Context) error {
	var newUser AppUser
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	query := "INSERT INTO appuser (household_id, username, email) VALUES (:household_id, :username, :email)"
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

// FetchUsers retrieves all users from the database.
func FetchItems(c echo.Context) error {
	if DB == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection is not established"})
	}

	var items []Items
	err := DB.Select(&items, "SELECT * FROM items")
	if err != nil {
		log.Printf("Error fetching items: %v", err)
		return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Error fetching items"})
	}

	return c.JSON(http.StatusOK, items)
}
