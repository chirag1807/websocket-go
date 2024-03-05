package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/chirag1807/websocket-go/api/model/dto"
	"github.com/chirag1807/websocket-go/api/model/request"
	"github.com/chirag1807/websocket-go/api/model/response"
	"github.com/chirag1807/websocket-go/api/service"
	"github.com/chirag1807/websocket-go/constants"
	errorhandling "github.com/chirag1807/websocket-go/error"
	"github.com/chirag1807/websocket-go/utils"
	"github.com/go-chi/chi/v5"
)

type ChatController interface {
	AddToSingleChat(w http.ResponseWriter, r *http.Request)
	GetAllChats(w http.ResponseWriter, r *http.Request)
	AddToRoomChat(w http.ResponseWriter, r *http.Request)
	CreateRoom(w http.ResponseWriter, r *http.Request)
	AddMemberToRoom(w http.ResponseWriter, r *http.Request)
	RemoveMemberFromRoom(w http.ResponseWriter, r *http.Request)
}

type chatController struct {
	chatService service.ChatService
}

func NewChatController(s service.ChatService) ChatController {
	return chatController{
		chatService: s,
	}
}

func (a chatController) AddToSingleChat(w http.ResponseWriter, r *http.Request) {
	var chat request.Chat

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadBodyError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &chat)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	err = a.chatService.AddToSingleChat(chat)

	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ChatNotAdded)
		return
	}

	response := response.SuccessResponse{
		Message: constants.CHAT_ADDED,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
}

func (a chatController) GetAllChats(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)

	chats, err := a.chatService.GetAllChats(id)

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	utils.ResponseGenerator(w, http.StatusOK, chats)
}

func (a chatController) AddToRoomChat(w http.ResponseWriter, r *http.Request) {
	var chat request.RoomChat

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadBodyError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &chat)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	err = a.chatService.AddToRoomChat(chat)

	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ChatNotAdded)
		return
	}

	response := response.SuccessResponse{
		Message: constants.CHAT_ADDED,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
}

func (a chatController) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var room request.RoomInfo

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadBodyError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &room)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	err = a.chatService.CreateRoom(room)

	if err != nil {
		log.Println(err)
		log.Println(room)
		utils.ErrorGenerator(w, errorhandling.RoomNotCreated)
		return
	}

	response := response.SuccessResponse{
		Message: constants.ROOM_CREATED,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
}

func (a chatController) AddMemberToRoom(w http.ResponseWriter, r *http.Request) {
	var room dto.RoomInfo

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadBodyError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &room)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	err = a.chatService.AddMemberToRoom(room)

	if err != nil {
		utils.ErrorGenerator(w, errorhandling.MemberNotAdded)
		return
	}

	response := response.SuccessResponse{
		Message: constants.MEMBER_ADDED,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
}

func (a chatController) RemoveMemberFromRoom(w http.ResponseWriter, r *http.Request) {
	var room dto.RoomInfo

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadBodyError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &room)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	err = a.chatService.RemoveMemberFromRoom(room)

	if err != nil {
		utils.ErrorGenerator(w, errorhandling.MemberNotRemoved)
		return
	}

	response := response.SuccessResponse{
		Message: constants.MEMBER_REMOVED,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
}
