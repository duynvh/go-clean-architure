package handlerimpl

import (
	"github.com/gin-gonic/gin"
	"go-clean-architure/handler"
	"go-clean-architure/libs/common"
	"go-clean-architure/models"
	"go-clean-architure/libs/utils"
	useCase "go-clean-architure/usecase"
	"net/http"
)

type UserHandlerImpl struct {
	useCase useCase.UserUseCase
}

func NewUserHandler(useCase useCase.UserUseCase) handler.UserHandler {
	return &UserHandlerImpl{
		useCase: useCase,
	}
}

func (h *UserHandlerImpl) Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	if email == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Please input all field"})
		return
	}

	user , _ := h.useCase.GetByEmail(email)
	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found"})
		return
	}

	if !utils.CheckHash(password, user.Password ) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Password is not correct"})
		return
	}

	// Generate token
	serialized := user.Serialize()
	token, _ := utils.GenerateToken(serialized)
	ctx.JSON(http.StatusCreated, common.JSON{
		"token": token,
	})
}

func (h *UserHandlerImpl) Register(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	if email == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Please input all field"})
		return
	}

	exists, _ := h.useCase.GetByEmail(email)
	if exists != nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "User is exist"})
		return
	}

	hash, hashErr := utils.Hash(password)
	if hashErr != nil {
		ctx.AbortWithStatus(500)
		return
	}

	user := models.User{Email: email, Password: hash}
	err := h.useCase.Create(&user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Errors"})
		return
	}

	// Generate token
	serialized := user.Serialize()
	token, _ := utils.GenerateToken(serialized)
	ctx.JSON(http.StatusCreated, common.JSON{
		"token": token,
	})
}
