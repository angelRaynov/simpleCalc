package repository

import (
	"calc/internal/calculator"
	"calc/internal/model"
	"sync"
)

type calculatorRepository struct {
	storage map[string]*model.Error
	mutex   *sync.RWMutex
}

func NewCalculatorRepository() calculator.StoreFinder {
	s := make(map[string]*model.Error)
	return &calculatorRepository{
		storage: s,
		mutex:   &sync.RWMutex{},
	}
}

func (cr *calculatorRepository) Store(e *model.Error) {
	cr.mutex.Lock()
	defer cr.mutex.Unlock()

	val, exist := cr.storage[e.Expression]
	if exist {
		val.Frequency++
		return
	}

	cr.storage[e.Expression] = e
	return
}

func (cr *calculatorRepository) FindAll() map[string]*model.Error {
	cr.mutex.RLock()
	defer cr.mutex.RUnlock()

	return cr.storage
}
