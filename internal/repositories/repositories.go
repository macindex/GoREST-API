package repositories

type Repositories struct {
	User interface {
		GetAll() []models.User 
		Add(newUser models.User)
	}
}

func New() *Repositories {
	return &Repositories()
}