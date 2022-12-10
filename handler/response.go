package handler

type CreateResponse struct {
	Id int `json:"id"`
}

type MakeFriendsResponse struct {
	Message string `json:"message"`
}
type DeleteFriendResponse struct {
	Message string `json:"message"`
}
