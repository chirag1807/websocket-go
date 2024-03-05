package service

import (
	"github.com/chirag1807/websocket-go/api/model/dto"
	"github.com/chirag1807/websocket-go/api/model/request"
	"github.com/chirag1807/websocket-go/api/model/response"
	"github.com/chirag1807/websocket-go/api/repository"
)

type ChatService interface {
	AddToSingleChat(chat request.Chat) error
	GetAllChats(id int64) ([]response.Chat, error)
	AddToRoomChat(chat request.RoomChat) error
	CreateRoom(room request.RoomInfo) error
	AddMemberToRoom(room dto.RoomInfo) error
	RemoveMemberFromRoom(room dto.RoomInfo) error
}

type chatService struct {
	chatRepository repository.ChatRepository
}

func NewChatService(a repository.ChatRepository) ChatService {
	return chatService{
		chatRepository: a,
	}
}

func (a chatService) AddToSingleChat(chat request.Chat) error {
	return a.chatRepository.AddToSingleChat(chat)
}

func (a chatService) GetAllChats(id int64) ([]response.Chat, error) {
	return a.chatRepository.GetAllChat(id)
}

func (a chatService) AddToRoomChat(chat request.RoomChat) error {
	return a.chatRepository.AddToRoomChat(chat)
}

func (a chatService) CreateRoom(room request.RoomInfo) error {
	return a.chatRepository.CreateRoom(room)
}

func (a chatService) AddMemberToRoom(room dto.RoomInfo) error {
	return a.chatRepository.AddMemberToRoom(room)
}

func (a chatService) RemoveMemberFromRoom(room dto.RoomInfo) error {
	return a.chatRepository.RemoveMemberFromRoom(room)
}
