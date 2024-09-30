package repository

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"murmur-server/model"
	"murmur-server/model/apperrors"
)

type friendRepository struct {
	DB *gorm.DB
}

func NewFriendRepository(db *gorm.DB) model.FriendRepository {
	return &friendRepository{
		DB: db,
	}
}

func (r *friendRepository) FriendsList(id string) (*[]model.Friend, error) {
	var friends []model.Friend

	result := r.DB.
		Table("users").
		Joins(`JOIN friends ON friends.user_id = "users".id`).
		Where("friends.friend_id = ?", id).
		Find(&friends)

	return &friends, result.Error
}

func (r *friendRepository) RequestList(id string) (*[]model.FriendRequest, error) {
	var requests []model.FriendRequest

	result := r.DB.
		Raw(`
		  select u.id, u.username, u.image, 1 as "type" from users u
		  join friend_requests fr on u.id = fr."sender_id"
		  where fr."receiver_id" = @id
		  UNION
		  select u.id, u.username, u.image, 0 as "type" from users u
		  join friend_requests fr on u.id = fr."receiver_id"
		  where fr."sender_id" = @id
		  order by username;
		`, sql.Named("id", id)).
		Find(&requests)

	return &requests, result.Error
}

func (r *friendRepository) FindByID(id string) (*model.User, error) {
	user := &model.User{}

	if err := r.DB.
		Preload("Friends").
		Preload("Requests").
		Where("id = ?", id).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, apperrors.NewNotFound("uid", id)
		}
		return user, apperrors.NewInternal()
	}

	return user, nil
}

func (r *friendRepository) DeleteRequest(memberId string, userId string) error {
	return r.DB.Exec(`
		DELETE
		FROM friend_requests
		WHERE receiver_id = @memberId AND sender_id = @userId
		   OR receiver_id = @userId AND sender_id = @memberId
`, sql.Named("memberId", memberId), sql.Named("userId", userId)).Error
}

func (r *friendRepository) RemoveFriend(memberId string, userId string) error {
	return r.DB.
		Exec("DELETE FROM friends WHERE user_id = ? AND friend_id = ?", memberId, userId).
		Exec("DELETE FROM friends WHERE user_id = ? AND friend_id = ?", userId, memberId).
		Error
}

func (r *friendRepository) Save(user *model.User) error {
	return r.DB.Save(&user).Error
}
