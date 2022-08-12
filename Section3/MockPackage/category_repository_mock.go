package MockPackage

import "github.com/stretchr/testify/mock"

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(Category)
		return &category
	}
}
