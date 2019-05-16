package handler

import "github.com/gin-gonic/gin"

type TodoHandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FetchById(ctx *gin.Context)
	Fetch(ctx *gin.Context)
}

