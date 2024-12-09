package usecase

import (
	"calc/internal/calculator"
	"calc/internal/model"
	"testing"
)

func TestCalculatorUseCase_EvaluateExpression(t *testing.T) {
	tests := []struct {
		calcRepo   calculator.StoreFinder
		expression string
		want       int
		wantErr    error
	}{
		{
			calcRepo:   nil,
			expression: "What is 5?",
			want:       5,
			wantErr:    nil,
		},
		{
			calcRepo:   nil,
			expression: "What is 5 plus 3?",
			want:       8,
			wantErr:    nil,
		},
		{
			calcRepo:   nil,
			expression: "What is 5 plus 3 multiplied by 6?",
			want:       48,
			wantErr:    nil,
		},
		{
			calcRepo:   nil,
			expression: "What is 5 plus 3 multiplied by 6 divided by 3?",
			want:       16,
			wantErr:    nil,
		},
		{
			calcRepo:   nil,
			expression: "What is 5 plus 3 multiplied by 6 divided by 3?",
			want:       16,
			wantErr:    nil,
		},
		{
			calcRepo:   nil,
			expression: "What is 5 plus 3 multiplied by 6 divided by 3 minus 2?",
			want:       14,
			wantErr:    nil,
		},
		{
			calcRepo:   nil,
			expression: "What is -5 plus -3?",
			want:       -8,
			wantErr:    nil,
		},
		{
			calcRepo:   nil,
			expression: "What is 0 plus -3?",
			want:       -3,
			wantErr:    nil,
		},
		{
			calcRepo:   nil,
			expression: "What is 3 times 3?",
			want:       9,
			wantErr:    nil,
		},
		{
			calcRepo:   nil,
			expression: "What is 10 divided by 0",
			want:       0,
			wantErr:    ErrZeroDivision,
		},
		{
			calcRepo:   nil,
			expression: "What is 10 divided",
			want:       0,
			wantErr:    ErrUnsupportedOperation,
		},
		{
			calcRepo:   nil,
			expression: "Who is Satoshi?",
			want:       0,
			wantErr:    ErrNonMathQuestion,
		},
		{
			calcRepo:   nil,
			expression: "What is 10 divided by plus minus 3",
			want:       0,
			wantErr:    ErrInvalidSyntax,
		},
		{
			calcRepo:   nil,
			expression: "",
			want:       0,
			wantErr:    ErrNonMathQuestion,
		},
		{
			calcRepo:   nil,
			expression: "What is 10 % 3",
			want:       0,
			wantErr:    ErrUnsupportedOperation,
		},
		{
			calcRepo:   nil,
			expression: "What is 10.222",
			want:       0,
			wantErr:    ErrInvalidSyntax,
		},
	}
	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			cu := &CalculatorUseCase{
				calcRepo: tt.calcRepo,
			}
			got, err := cu.EvaluateExpression(tt.expression)
			if err != tt.wantErr {
				t.Errorf("EvaluateExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EvaluateExpression() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type MockStore struct {
	calculator.Storer
}

func (n *MockStore) FindAll() map[string]*model.Error {
	m := make(map[string]*model.Error)
	m["test"] = &model.Error{}
	m["test2"] = &model.Error{}
	return m
}

func (n *MockStore) Store(e *model.Error) {
}

func TestCalculatorUseCase_FindErrors(t *testing.T) {
	tests := struct {
		name         string
		calcRepo     calculator.StoreFinder
		wantErrCount int
	}{
		name:         "find errors",
		calcRepo:     &MockStore{},
		wantErrCount: 2,
	}
	t.Run(tests.name, func(t *testing.T) {
		cu := CalculatorUseCase{
			calcRepo: tests.calcRepo,
		}
		if got := cu.FindErrors(); len(got) != tests.wantErrCount {
			t.Errorf("FindErrors() = %v, want %v", got, tests.wantErrCount)
		}
	})
}

func TestCalculatorUseCase_StoreError(t *testing.T) {
	tests := struct {
		name     string
		calcRepo calculator.StoreFinder
		want     []*model.Error
		arg      model.Error
	}{
		name:     "store error",
		calcRepo: &MockStore{},
		want:     nil,
		arg: model.Error{
			Expression: "test",
			Endpoint:   "/eval",
			Frequency:  1,
			Type:       "test",
		},
	}
	t.Run(tests.name, func(t *testing.T) {
		cu := CalculatorUseCase{
			calcRepo: tests.calcRepo,
		}
		cu.StoreError(&tests.arg)
	})
}
