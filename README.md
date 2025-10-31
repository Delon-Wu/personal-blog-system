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
│   │   └── post.go            # Post model definition
│   ├── controllers            # Request handlers
│   │   └── auth_controller.go # Auth-related request handling
│   │   └── user_controller.go # User-related request handling
│   │   └── base_controller.go # Base-related request handling
│   │   └── post_controller.go # Post-related request handling
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

### 准备

- Go (version 1.16 or later)
- Gin
- GORM
- A database (MySQL) （创建mysql表，表名：personal_blog_system）
- .env （参照example.env配置环境变量文件）

### 依赖安装

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

### 运行

To run the application, execute the following command:

```
go run src/main.go
```

The server will start on `http://localhost:8080`.

### API Endpoints

`POST   /api/login`                登录  
`POST   /api/users`                注册  
`GET    /api/users/:id`            查询用户信息  
`POST   /api/post`                 提交文章  
`PUT    /api/post`                 修改文章  
`GET    /api/post/:id`             获取文章信息
`POST   /api/post/comment/:id`     创建评论
`GET    /api/post/comment/:id`     获取评论信息  
`DELETE /api/post/comment/:id`     删除评论  
`GET    /api/post/list`            获取文章列表  
`DELETE /api/post/:id`             删除文章  

### License

This project is licensed under the MIT License.