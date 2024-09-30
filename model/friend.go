package model

type Friend struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Image    string `json:"image"`
	IsOnline bool   `json:"isOnline"`
}
type FriendService interface {
	GetFriends(id string) (*[]Friend, error)
	GetRequests(id string) (*[]FriendRequest, error)
	GetMemberById(id string) (*User, error)
	DeleteRequest(memberId string, userId string) error
	RemoveFriend(memberId string, userId string) error
	SaveRequests(user *User) error
}

type FriendRepository interface {
	FindByID(id string) (*User, error)
	FriendsList(id string) (*[]Friend, error)
	RequestList(id string) (*[]FriendRequest, error)
	DeleteRequest(memberId string, userId string) error
	RemoveFriend(memberId string, userId string) error
	Save(user *User) error
}
