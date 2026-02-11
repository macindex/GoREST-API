package usecases


type UseCases struct {
	repos *repositories.Repositories
}

func New(repos *repositories.Repositories) *UseCases{
	return &UseCases {
		repos: repos,
	}
}

func (u UseCases) GetAll() ([]models.User, error) {
	u.repos.User.GetAll()
} 
func (u UseCases) Add(newUser models.User) {

}
