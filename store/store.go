package store

import (
	"fmt"
)

type UserStore interface {
	Put(u UserInterface) (id int)
	GetAll() (res string)
	GetUserById(id int) (UserInterface, error)
	DeleteUser(id int)
}

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

type Store struct {
	users map[int]UserInterface
}

func NewStore() *Store {
	return &Store{users: make(map[int]UserInterface)}
}

func (s *Store) nextId() (res int) {
	for id, _ := range s.users {
		if id > res {
			res = id
		}
	}
	return res + 1
}

func (s *Store) Put(u UserInterface) (id int) {
	id = s.nextId()
	u.SetId(id)
	s.users[id] = u
	return id
}

func (s *Store) GetAll() (res string) {
	for _, u := range s.users {
		res += u.ToString()
	}
	return
}

func (s *Store) GetUserById(id int) (UserInterface, error) {
	u, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user id: %d not found", id)
	}
	return u, nil
}

func (s *Store) DeleteUser(id int) {
	delete(s.users, id)
}
