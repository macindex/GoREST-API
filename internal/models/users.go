package models

type User struct {
	users []models.User
}

type CreateUSerRequest struct {
	Name string 
}

func New() *Users {
	return &Users {users: make([]models.User, 0)}
}

func (u Users) GetAll() []models.User {
	return u.users
}

func (u Users) EmailExists() (bool) {
	for _, v range u.users {
		if v.Email == email {
			return true
		}
	}
	return false
}

func (u *Users) Add(newUser models.User) {
	u.users = append(u.users, newUser)
}
