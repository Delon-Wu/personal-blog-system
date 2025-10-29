# Personal Blog System

This project is a personal blog system built using Go, Gin, and GORM. It provides a simple API for managing users and their blog posts.

## Project Structure

```
personal-blog-system
├── src
│   ├── main.go                # Entry point of the application
│   ├── config                 # Configuration files
│   │   └── config.go          # Loads and manages configuration
│   ├── database               # Database connection and operations
│   │   └── gorm.go            # Initializes GORM database connection
│   ├── models                 # Data models
│   │   └── user.go            # User model definition
│   ├── controllers            # Request handlers
│   │   └── user_controller.go  # User-related request handling
│   ├── routes                 # API routes
│   │   └── routes.go          # Sets up application routes
│   └── migrations             # Database migrations
│       └── 0001_create_users.sql # SQL for creating users table
├── go.mod                     # Go module configuration
├── .env                       # Environment variables
├── .gitignore                 # Files and directories to ignore in version control
└── README.md                  # Project documentation and usage instructions
```

## Getting Started

### Prerequisites

- Go (version 1.16 or later)
- Gin
- GORM
- A database (e.g., MySQL, PostgresSQL)

### Installation

1. Clone the repository:

   ```
   git clone <repository-url>
   cd personal-blog-system
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Set up your database and update the `.env` file with your database connection string.

### Running the Application

To run the application, execute the following command:

```
go run src/main.go
```

The server will start on `http://localhost:8080`.

### API Endpoints

- `POST /users` - Create a new user
- `GET /users/:id` - Get user information by ID

### License

This project is licensed under the MIT License.