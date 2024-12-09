package repository

import (
	"calc/internal/model"
	"sync"
	"testing"
)

func Test_calculatorRepository_Store(t *testing.T) {
	type fields struct {
		storage map[string]*model.Error
		mutex   *sync.RWMutex
	}

	tests := []struct {
		name              string
		fields            fields
		args              *model.Error
		expectedErrCount  int
		expectedFrequency int
	}{
		{
			name: "store error",
			fields: fields{
				storage: map[string]*model.Error{},
				mutex:   &sync.RWMutex{},
			},
			args:              &model.Error{Expression: "test", Frequency: 1},
			expectedErrCount:  1,
			expectedFrequency: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := &calculatorRepository{
				storage: tt.fields.storage,
				mutex:   tt.fields.mutex,
			}

			cr.Store(tt.args)
			cr.Store(tt.args)
			errs := cr.FindAll()
			errCount := len(errs)
			if errCount != tt.expectedErrCount {
				t.Errorf("Expected %d errors, got %d", tt.expectedErrCount, errCount)
			}

			val, exist := errs["test"]
			if exist {
				if val.Frequency != tt.expectedFrequency {
					t.Errorf("Expected frequency %d , got %d", tt.expectedFrequency, val.Frequency)
				}
			}
		})
	}
}
