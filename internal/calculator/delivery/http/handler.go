package http

import (
	"calc/internal/calculator"
	"calc/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CalculatorHandler struct {
	calcUseCase calculator.UseCaser
}

func NewCalculatorHandler(cu calculator.UseCaser) calculator.Handler {
	return &CalculatorHandler{
		calcUseCase: cu,
	}
}

func (ch *CalculatorHandler) EvaluateHandler(c *gin.Context) {
	var req model.EvaluateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	result, err := ch.calcUseCase.EvaluateExpression(req.Expression)
	if err != nil {
		e := model.Error{
			Expression: req.Expression,
			Endpoint:   "/evaluate",
			Frequency:  1,
			Type:       err.Error(),
		}
		ch.calcUseCase.StoreError(&e)

		c.IndentedJSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, model.EvaluateResponse{Result: result})
	return
}

func (ch *CalculatorHandler) ValidateHandler(c *gin.Context) {
	var req model.ValidateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	_, err := ch.calcUseCase.EvaluateExpression(req.Expression)
	if err != nil {
		e := model.Error{
			Expression: req.Expression,
			Endpoint:   "/validate",
			Frequency:  1,
			Type:       err.Error(),
		}
		ch.calcUseCase.StoreError(&e)

		c.IndentedJSON(http.StatusOK, model.ValidateResponse{Valid: false})
		return
	}

	c.IndentedJSON(http.StatusOK, model.ValidateResponse{Valid: true})
	return
}

func (ch *CalculatorHandler) ErrorHandler(c *gin.Context) {
	errs := ch.calcUseCase.FindErrors()
	if len(errs) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no records"})
		return
	}
	c.IndentedJSON(http.StatusOK, errs)
}
