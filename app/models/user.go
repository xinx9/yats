package models

type UserStorage interface {
	List() []*User
	Get(int) *User
	Update(int, User) *User
	Create(User)
	Delete(int) *User
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserStore struct {
}

var users = []*User{
	{
		ID:   1,
		Name: "John Doe",
	},
	{
		ID:   2,
		Name: "Jane Smith",
	},
}

func (u UserStore) Get(id int) *User {
	for _, user := range users {
		if user.ID == id {
			return user
		}
	}
	return nil
}

func (u UserStore) List() []*User {
	return users
}

func (u UserStore) Create(user User) {
	users = append(users, &user)
}

func (u UserStore) Delete(id int) *User {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], (users)[i+1:]...)
			return &User{}
		}
	}
	return nil
}

func (u UserStore) Update(id int, userUpdate User) *User {
	for i, user := range users {
		if user.ID == id {
			users[i] = &userUpdate
			return users[i]
		}
	}
	return nil
}
