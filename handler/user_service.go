package handler

type UserService interface {
	CreateUser(name string, age int) int
	GetAllUsers() string
	MakeFriends(id1, id2 int) (name1, name2 string, err error)
	DeleteUser(id int) (string, error)
	GetUserFriends(id int) (res string, err error)
	UpdateAge(id, age int) error
}
