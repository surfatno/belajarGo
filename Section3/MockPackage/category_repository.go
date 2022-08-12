package MockPackage

type CategoryRepository interface {
	FindById(id string) *Category
}
