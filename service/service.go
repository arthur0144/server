package service

import (
	"server/store"
)

type Service struct {
	store store.Store
}

func NewService(s store.Store) Service {
	return Service{
		store: s,
	}
}

func (s *Service) CreateUser(name string, age int) int {
	u := store.NewUser(name, age)
	return s.store.Put(u)
}

func (s *Service) GetAllUsers() string {
	return s.store.GetAll()
}

func (s *Service) MakeFriends(id1, id2 int) (name1, name2 string, err error) {
	user1, err := s.store.GetUserById(id1)
	if err != nil {
		return "", "", err
	}
	user2, err := s.store.GetUserById(id2)
	if err != nil {
		return "", "", err
	}
	user1.MakeFriends(id2)
	user2.MakeFriends(id1)
	return user1.Name(), user2.Name(), nil
}

func (s *Service) DelUser(idUs int) (string, error) {
	user, err := s.store.GetUserById(idUs)
	if err != nil {
		return "", err
	}
	defer s.store.DeleteUser(idUs)
	friends := user.GetFriends()
	for _, u := range friends {
		user, _ := s.store.GetUserById(u)
		err := user.DeleteFriend(idUs)
		if err != nil {
			return "", err
		}
	}
	return user.Name(), nil
}
