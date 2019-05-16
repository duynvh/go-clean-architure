package handlerimpl

import (
	"github.com/gin-gonic/gin"
	"go-clean-architure/handler"
	"go-clean-architure/libs/common"
	"go-clean-architure/models"
	useCase "go-clean-architure/usecase"
	"net/http"
	"strconv"
)

type TodoHandlerImpl struct {
	useCase useCase.TodoUseCase
}

func NewTodoHandler(useCase useCase.TodoUseCase) handler.TodoHandler {
	return &TodoHandlerImpl{
		useCase: useCase,
	}
}

func (h *TodoHandlerImpl) Create(ctx *gin.Context) {
	title := ctx.PostForm("title")
	complete := ctx.PostForm("completed")

	if title == "" || complete == ""  {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Please input all field"})
		return
	}

	completed, _ := strconv.Atoi(complete)
	user := ctx.MustGet("user").(models.User)
	todo := models.Todo{Title: title, Completed: completed, UserID: user.ID}
	err := h.useCase.Create(&todo)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Errors"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

func (h *TodoHandlerImpl) Fetch(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	todos, _ := h.useCase.Fetch(page)

	if len(todos) <= 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	length := len(todos)
	serialized := make([]common.JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = todos[i].Serialize()
	}

	ctx.JSON(http.StatusOK, common.JSON{
		"data": serialized,
		"page": page,
		"status": http.StatusOK,
	})
}

func (h *TodoHandlerImpl) Update(ctx *gin.Context) {
	title := ctx.PostForm("title")
	complete := ctx.PostForm("completed")

	if title == "" || complete == ""  {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Please input all field"})
		return
	}

	completed, _ := strconv.Atoi(complete)

	id,_ := strconv.Atoi(ctx.Param("id"))
	user := ctx.MustGet("user").(models.User)

	todo, _ := h.useCase.FetchById(id)
	if todo == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not Found"})
		return
	}

	if todo.UserID != user.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "Not Forbidden!"})
		return
	}

	todo.Title = title
	todo.Completed = completed
	h.useCase.Update(todo)
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

func (h *TodoHandlerImpl) Delete(ctx *gin.Context) {
	id,_ := strconv.Atoi(ctx.Param("id"))
	user := ctx.MustGet("user").(models.User)

	todo, _ := h.useCase.FetchById(id)
	if todo == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not Found"})
		return
	}

	if todo.UserID != user.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "Not Forbidden!"})
		return
	}

	h.useCase.Delete(id)
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}

func (h *TodoHandlerImpl) FetchById(ctx *gin.Context) {
	id,_ := strconv.Atoi(ctx.Param("id"))
	todo, _ := h.useCase.FetchById(id)

	if todo == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	ctx.JSON(http.StatusOK, common.JSON{
		"data": todo.Serialize(),
		"status": http.StatusOK,
	})
}

