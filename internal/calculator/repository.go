package calculator

import "calc/internal/model"

type StoreFinder interface {
	Storer
	Finder
}

type Finder interface {
	FindAll() map[string]*model.Error
}
type Storer interface {
	Store(e *model.Error)
}
