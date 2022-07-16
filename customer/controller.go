package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	InputCustomerDataHandler(*gin.Context)
	FindAllPointsHandler(*gin.Context)
}

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{service}
}

// swagger:route POST /api/customer customer inputCustomerData
// Input new customer data
//
// responses:
//		200: inputCustomerData
//		422: messageResponse
//		500: messageResponse
func (c *controller) InputCustomerDataHandler(ctx *gin.Context) {
	var customer Customer

	err := ctx.ShouldBindJSON(&customer)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	customer, err = c.service.Save(customer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

// swagger:route GET /api/point customer findAllPoints
// Find points of all customer
//
// responses:
//		200: findAllPoints
//		500: messageResponse
func (c *controller) FindAllPointsHandler(ctx *gin.Context) {
	points, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, points)
}
