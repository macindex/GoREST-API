package repositories

type Repositories struct {
	User interface {
		GetAll() []models.User 
		Add(newUser models.User)
		EmailExists(email string) bool
	}
}

func New() *Repositories {
	return &Repositories()
}
