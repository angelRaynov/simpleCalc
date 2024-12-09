package calculator

import "calc/internal/model"

type UseCaser interface {
	EvaluateExpression(expression string) (int, error)
	StoreError(e *model.Error)
	FindErrors() []*model.Error
}
