package models

type UserStorage interface {
	List() []*User
	Get(int) *User
	Update(int, string) *User
	Create(User)
	Delete(int) *User
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Role      string `json:"Role"`
	Status    bool   `json:"Status"`
}

type UserStore struct {
}

var users = []*User{
	{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Role:      "Admin",
		Status:    true,
	},
	{
		ID:        2,
		FirstName: "jane",
		LastName:  "smith",
		Role:      "Admin",
		Status:    true,
	},
	{
		ID:        3,
		FirstName: "Joe",
		LastName:  "Dirst",
		Role:      "User",
		Status:    true,
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

func (u UserStore) Update(id int, userName string) *User {
	for i, user := range users {
		if user.ID == id {
			users[i].FirstName = userName
			return users[i]
		}
	}
	return nil
}
