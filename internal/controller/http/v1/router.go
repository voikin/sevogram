package httpv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(engine *gin.Engine) {
	baseRoute := engine.Group("/v1")

	baseRoute.GET("ping", ping)
}

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}
