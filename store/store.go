package store

import (
	"fmt"
)

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

func (u *User) toString() string {
	return fmt.Sprintf("Name is %s , Age %d is , friends %d and Id: %d\n", u.name, u.age, u.friends, u.id)
}

func (u *User) Name() string {
	return u.name
}

func (u *User) GetFriends() []int {
	return u.friends
}

func (u *User) MakeFriends(idUs int) {
	u.friends = append(u.friends, idUs)
}

func (u *User) DeleteFriend(idF int) error {
	for i, us := range u.friends {
		if us == idF {
			u.friends = append(u.friends[:i], u.friends[i+1:len(u.friends)]...)
			return nil
		}
	}
	return fmt.Errorf("user id: %d not found", idF)
}

type Store map[int]*User

func (s Store) Put(u *User) int {
	u.id = len(s) + 1
	s[u.id] = u
	return u.id
}

func (s *Store) GetAll() (res string) {
	for _, user := range *s {
		res += user.toString()
	}
	return
}

func (s *Store) GetUserById(id int) (*User, error) {
	for idUs, user := range *s {
		if id == idUs {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user id: %d not found", id)
}

func (s *Store) DeleteUser(id int) {
	delete(*s, id)
}
