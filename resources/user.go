package resources

import (
	"sync"

	"go-rest/utils"
)

type UserContainer struct {
	lock  sync.RWMutex
	users map[string]*User
}

func (uc *UserContainer) AllUsers() []*User {
	uc.lock.RLock()
	defer uc.lock.RUnlock()

	users := make([]*User, 0, len(uc.users))
	for _, user := range uc.users {
		users = append(users, user)
	}
	return users
}

func (uc *UserContainer) Add(user *User) {
	uc.lock.Lock()
	defer uc.lock.Unlock()

	uc.users[user.Id] = user
}

func (uc *UserContainer) Remove(id string) {
	uc.lock.Lock()
	defer uc.lock.Unlock()

	delete(uc.users, id)
}

var userContainer *UserContainer

func init() {
	userContainer = &UserContainer{
		lock:  sync.RWMutex{},
		users: make(map[string]*User),
	}
}

type User struct {
	Id    string `json:"string"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func NewUser(name string, age int, email string) *User {
	return &User{
		Id:    utils.RandomString(10),
		Name:  name,
		Age:   age,
		Email: email,
	}
}

func AddUser(name string, age int, email string) *User {
	user := NewUser(name, age, email)
	userContainer.Add(user)
	return user
}

func RemoveUser(id string) {
	userContainer.Remove(id)
}

func ListUsers() []*User {
	return userContainer.AllUsers()
}
