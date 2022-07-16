package passbook

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	FindPassbooksByAccountIdHandler(*gin.Context)
}

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{service}
}

// swagger:route GET /api/passbook/{account_id}/{start_date}/{end_date} passbook findPassbooksCustomer
// Find customer passbooks
//
// responses:
//		200: findPassbooksCustomer
//		400: messageResponse
//		500: messageResponse
func (c *controller) FindPassbooksByAccountIdHandler(ctx *gin.Context) {
	startDate := ctx.Params.ByName("start_date")
	endDate := ctx.Params.ByName("end_date")
	accountIdString := ctx.Params.ByName("account_id")
	accountId, err := strconv.Atoi(accountIdString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	passbooks, err := c.service.FindByAccountId(uint(accountId), startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, passbooks)
}
