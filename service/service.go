package service

import (
	"log"

	"server/store"
)

type Service struct {
	store store.Store
}

func NewService() Service {
	return Service{
		store: store.NewStore(),
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
	user1.AddFriend(id2)
	user2.AddFriend(id1)
	return user1.Name(), user2.Name(), nil
}

func (s *Service) DeleteUser(id int) (string, error) {
	user, err := s.store.GetUserById(id)
	if err != nil {
		return "", err
	}
	s.store.DeleteUser(id)

	friends := user.GetFriends()
	for _, fid := range friends {
		u, err := s.store.GetUserById(fid)
		if err != nil {
			log.Printf("can't get friend id=%d of user with id=%d", fid, id)
			continue
		}
		u.DeleteFriend(id)
	}

	return user.Name(), nil
}

func (s *Service) GetAllUserFriends(id int) (res string, err error) {
	user, err := s.store.GetUserById(id)
	if err != nil {
		return "", err
	}
	friends := user.GetFriends()
	for _, fid := range friends {
		u, err := s.store.GetUserById(fid)
		if err != nil {
			log.Printf("can't get friend id=%d of user with id=%d", fid, id)
			continue
		}
		res += u.ToString()
	}
	return
}
