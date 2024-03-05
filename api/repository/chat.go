package repository

import (
	"context"
	"fmt"

	"github.com/chirag1807/websocket-go/api/model/dto"
	"github.com/chirag1807/websocket-go/api/model/request"
	"github.com/chirag1807/websocket-go/api/model/response"
	"github.com/jackc/pgx/v5"
)

type ChatRepository interface {
	AddToSingleChat(chat request.Chat) error
	GetAllChat(id int64) ([]response.Chat, error)
	AddToRoomChat(chat request.RoomChat) error
	CreateRoom(room request.RoomInfo) error
	AddMemberToRoom(room dto.RoomInfo) error
	RemoveMemberFromRoom(room dto.RoomInfo) error
}

type chatRepository struct {
	pgx *pgx.Conn
}

func NewChatRepo(pgx *pgx.Conn) ChatRepository {
	return chatRepository{
		pgx: pgx,
	}
}

func (a chatRepository) AddToSingleChat(chat request.Chat) error {
	var chatID int64
	err := a.pgx.QueryRow(context.Background(), `INSERT INTO singlechat (senderid, receiverid, message) VALUES ($1, $2, $3) RETURNING id`, chat.Sender, chat.Receiver, chat.Message).Scan(&chatID)
	if err != nil {
		return err
	}
	return nil
}

func (a chatRepository) GetAllChat(id int64) ([]response.Chat, error) {
	chats, err := a.pgx.Query(context.Background(), `SELECT * FROM singlechat WHERE senderid = $1 or receiverid = $1`, id)
	chatSlice := make([]response.Chat, 0)

	if err != nil {
		return chatSlice, err
	}
	defer chats.Close()

	var chat response.Chat
	for chats.Next() {
		if err := chats.Scan(&chat.ID, &chat.Sender, &chat.Receiver, &chat.Message, &chat.Time); err != nil {
			return chatSlice, err
		}
		chatSlice = append(chatSlice, chat)
	}

	return chatSlice, nil
}

func (a chatRepository) AddToRoomChat(chat request.RoomChat) error {
	var chatID int64
	err := a.pgx.QueryRow(context.Background(), `INSERT INTO roomchat (roomid, senderid, message) VALUES ($1, $2, $3) RETURNING id`, chat.RoomID, chat.Sender, chat.Message).Scan(&chatID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (a chatRepository) CreateRoom(room request.RoomInfo) error {
	_, err := a.pgx.Exec(context.Background(), `INSERT INTO roominfo (rooname, createdby, members) VALUES ($1, $2, $3)`, room.RoomName, room.CreatedBy, room.Members)
	if err != nil {
		return err
	}
	return nil
}

func (a chatRepository) AddMemberToRoom(room dto.RoomInfo) error {
	// _, err := a.pgx.Exec(context.Background(), `UPDATE roominfo SET members = ARRAY_APPEND(members, $1) WHERE id = $2`, room.Members, room.ID)
	_, err := a.pgx.Exec(context.Background(), `UPDATE roominfo SET members = $1 WHERE id = $2`, room.Members, room.ID)
	if err != nil {
		return err
	}
	return nil
}

func (a chatRepository) RemoveMemberFromRoom(room dto.RoomInfo) error {
	_, err := a.pgx.Exec(context.Background(), `UPDATE roominfo SET members = ARRAY_REMOVE(members, $1) WHERE id = $2`, room.Members[0], room.ID)
	if err != nil {
		return err
	}
	return nil
}
