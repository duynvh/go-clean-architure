package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"go-clean-architure/driver"
	"go-clean-architure/handler/handlerimpl"
	"go-clean-architure/libs/middlewares"
	"go-clean-architure/models"
	"go-clean-architure/repositories/repoimpl"
	"go-clean-architure/usecase/usecaseimpl"
)

func LoadEnviroment() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
}

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func RunApp() {
	LoadEnviroment()

	db := driver.Connect()

	db.SQL.AutoMigrate(&models.User{})
	db.SQL.AutoMigrate(&models.Todo{})
	// Config App router
	router := gin.Default()
	Router(router, db.SQL)

	router.Run()
}

// Init Router

func Router(router *gin.Engine, db *gorm.DB) {
	router.Use(CORSMiddleware())
	router.Use(middlewares.JWTMiddleware())
	// User
	repoUser := repoimpl.NewUserRepo(db)
	useCaseUser := usecaseimpl.NewUserUsecase(repoUser)
	handlerUser := handlerimpl.NewUserHandler(useCaseUser)

	// Todo
	repoTodo := repoimpl.NewTodoRepo(db)
	useCaseTodo := usecaseimpl.NewTodoUsecase(repoTodo)
	handlerTodo := handlerimpl.NewTodoHandler(useCaseTodo)

	v1 := router.Group("api/v1")
	{
		v1.POST("/user/register", handlerUser.Register)
		v1.POST("/user/login", handlerUser.Login)

		v1.GET("/todo", handlerTodo.Fetch)
		v1.POST("/todo", middlewares.Authorized, handlerTodo.Create)
		v1.DELETE("/todo/:id", middlewares.Authorized, handlerTodo.Delete)
		v1.PUT("/todo/:id", middlewares.Authorized, handlerTodo.Update)
		v1.GET("/todo/:id", handlerTodo.FetchById)
	}
}

func main() {
	RunApp()
}
