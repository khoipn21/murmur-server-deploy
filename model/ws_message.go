package model

import (
	"encoding/json"
	"log"
)

type ReceivedMessage struct {
	Action  string `json:"action"`
	Room    string `json:"room"`
	Message *any   `json:"message"`
}

type WebsocketMessage struct {
	Action string `json:"action"`
	Data   any    `json:"data"`
}

func (message *WebsocketMessage) Encode() []byte {
	encoding, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	return encoding
}

type SocketService interface {
	EmitNewMessage(room string, message *MessageResponse)
	EmitEditMessage(room string, message *MessageResponse)
	EmitDeleteMessage(room, messageId string)

	EmitNewChannel(room string, channel *ChannelResponse)
	EmitNewPrivateChannel(members []string, channel *ChannelResponse)
	EmitEditChannel(room string, channel *ChannelResponse)
	EmitDeleteChannel(channel *Channel)

	EmitEditGuild(guild *Guild)
	EmitDeleteGuild(guildId string, members []string)
	EmitRemoveFromGuild(memberId, guildId string)

	EmitAddMember(room string, member *User)
	EmitRemoveMember(room, memberId string)

	EmitNewDMNotification(channelId string, user *User)
	EmitNewNotification(guildId, channelId string)

	EmitSendRequest(room string)
	EmitAddFriendRequest(room string, request *FriendRequest)
	EmitAddFriend(user, member *User)
	EmitRemoveFriend(userId, memberId string)
}
