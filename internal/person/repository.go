package person

type Repository interface {
	FindById(id int) (Person, error)
	FindByTerm(t string) ([]Person, error)
	Count() (int, error)
	Insert(p Person) (int, error)
}
