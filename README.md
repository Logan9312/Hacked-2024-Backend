# Hacked-2024-Backend

# Description

I built this backend to support the roomates App, it was made using the following technologies :

- Go
- SQLX
- Postgres
- Echo Router

# How to run

- Clone the repo
- Create a `.env` file in the root directory, and fill out variables based on `.env.example`
- Run `go run .` in the root directory
- The server will be running on port 8080

# Endpoints

/Health - GET - Health check endpoint
/payments - GET - Fetches all payments
/users - GET - Fetches all users
/tasks - GET - Fetches all tasks
/lists - GET - Fetches all lists
/messages - GET - Fetches all messages
/items - GET - Fetches all items

/save/payment - POST - Saves a payment
/save/user - POST - Saves a user
/save/tasks - POST - Saves a task
/save/lists - POST - Saves a list
/save/messages - POST - Saves a message
