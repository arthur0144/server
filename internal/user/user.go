package user

import (
	"fmt"
)

type UserInterface interface {
	ToString() string
	Name() string
	Age(a int)
	GetFriends() []int
	AddFriend(id int)
	DeleteFriend(id int)
	SetId(id int)
}

type User struct {
	id      int
	name    string
	age     int
	friends []int
}

func NewUser(name string, age int) *User {
	return &User{
		name: name,
		age:  age,
	}
}

func (u *User) ToString() string {
	return fmt.Sprintf("Name is %s , Age %d is , friends %d and Id: %d", u.name, u.age, u.friends, u.id)
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Age(a int) {
	u.age = a
}

func (u *User) GetFriends() []int {
	return u.friends
}

func (u *User) AddFriend(id int) {
	u.friends = append(u.friends, id)
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u *User) DeleteFriend(id int) {
	for i := range u.friends {
		if u.friends[i] == id {
			u.friends = append(u.friends[:i], u.friends[i+1:len(u.friends)]...)
			return
		}
	}
	return
}
