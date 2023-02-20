package helper

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleResponse (ctx *gin.Context, res any, err error) {
	if err != nil {
		handleResponseError(ctx, err)
	} else {
		ctx.JSON(http.StatusOK, res)
	}
}

func handleResponseError (ctx *gin.Context, err error) {
	if (err.Error() == "Something went wrong") {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}