package store

import (
	"fmt"

	"server/internal/user"
)

type StoreInterface interface {
	Put(u user.UserInterface) (id int)
	GetAll() (res string)
	GetUserById(id int) (user.UserInterface, error)
	DeleteUser(id int)
}

type Store struct {
	users map[int]user.UserInterface
}

func NewStore() *Store {
	return &Store{users: make(map[int]user.UserInterface)}
}

func (s *Store) nextId() (res int) {
	for id, _ := range s.users {
		if id > res {
			res = id
		}
	}
	return res + 1
}

func (s *Store) Put(u user.UserInterface) (id int) {
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

func (s *Store) GetUserById(id int) (user.UserInterface, error) {
	u, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user id: %d not found", id)
	}
	return u, nil
}

func (s *Store) DeleteUser(id int) {
	delete(s.users, id)
}
