package MockPackage

import "errors"

type CategoryService struct {
	Repository CategoryRepository
}

func (service CategoryService) Get(id string) (*Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
