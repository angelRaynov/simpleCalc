package usecase

import (
	"calc/internal/calculator"
	"calc/internal/model"
	"errors"
	"strconv"
	"strings"
)

var (
	ErrNonMathQuestion      = errors.New("non math question")
	ErrUnsupportedOperation = errors.New("unsupported operation")
	ErrInvalidSyntax        = errors.New("invalid syntax")
	ErrZeroDivision         = errors.New("division by zero")
)

type CalculatorUseCase struct {
	calcRepo calculator.StoreFinder
}

func NewCalculatorUseCase(calcRepo calculator.StoreFinder) *CalculatorUseCase {
	return &CalculatorUseCase{
		calcRepo: calcRepo,
	}
}

func (cu *CalculatorUseCase) EvaluateExpression(expression string) (int, error) {
	expression = strings.ToLower(expression)
	prefix := "what is"
	// Check if it's a non-math question
	if !strings.HasPrefix(expression, prefix) {
		return 0, ErrNonMathQuestion
	}

	// Remove "what is" prefix
	expression = strings.TrimPrefix(expression, prefix)
	expression = strings.TrimSuffix(expression, "?")

	// Remove "by"
	expression = strings.ReplaceAll(expression, "by", "")

	operands := strings.Fields(expression)
	operandsCount := len(operands)

	if operandsCount == 2 || len(operands)%2 == 0 {
		return 0, ErrUnsupportedOperation
	}

	result, err := strconv.Atoi(operands[0])
	if err != nil {
		return 0, ErrInvalidSyntax
	}

	if operandsCount == 1 {
		return result, nil
	}

	for i := 1; i < operandsCount-1; i += 2 {
		operator := operands[i]
		operand, err := strconv.Atoi(operands[i+1])
		if err != nil {
			return 0, ErrInvalidSyntax
		}

		switch operator {
		case "plus":
			result += operand
		case "minus":
			result -= operand
		case "multiplied", "times":
			result *= operand
		case "divided":
			if operand == 0 {
				return 0, ErrZeroDivision
			}
			result /= operand
		default:
			return 0, ErrUnsupportedOperation
		}
	}

	return result, nil
}

func (cu *CalculatorUseCase) StoreError(e *model.Error) {
	cu.calcRepo.Store(e)
}

func (cu *CalculatorUseCase) FindErrors() []*model.Error {
	errs := cu.calcRepo.FindAll()

	var listErr []*model.Error

	for _, v := range errs {
		listErr = append(listErr, v)
	}

	return listErr
}
