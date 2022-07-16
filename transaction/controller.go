package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	InputTransaction(*gin.Context)
}

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{service}
}

// swagger:route POST /api/transaction transaction inputTransaction
// Inpit transaction
//
// responses:
//		200: messageResponse
//		422: messageResponse
//		500: messageResponse
func (c *controller) InputTransaction(ctx *gin.Context) {
	var transaction Transaction

	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	err = c.service.Save(transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
