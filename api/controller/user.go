package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/chirag1807/websocket-go/api/model/request"
	"github.com/chirag1807/websocket-go/api/model/response"
	"github.com/chirag1807/websocket-go/api/service"
	"github.com/chirag1807/websocket-go/api/validation"
	"github.com/chirag1807/websocket-go/constants"
	"github.com/chirag1807/websocket-go/error"
	"github.com/chirag1807/websocket-go/utils"
)

type AuthController interface {
	UserRegistration(w http.ResponseWriter, r *http.Request)
	UserLogin(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(s service.AuthService) AuthController {
	return authController{
		authService: s,
	}
}

func (a authController) UserRegistration(w http.ResponseWriter, r *http.Request) {
	var user request.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadBodyError)
		return
	}
	defer r.Body.Close()

	// err := json.NewDecoder(r.Body).Decode(&user)

	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	isEmail := validation.EmailValidation(user.Email)
	if !isEmail {
		utils.ErrorGenerator(w, errorhandling.EmailvalidationError)
		return
	}

	err = a.authService.UserRegistration(user)
	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.USER_REGISTRATION_SUCCEED,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
}

func (a authController) UserLogin(w http.ResponseWriter, r *http.Request) {
	var userLoginRequest request.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadBodyError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &userLoginRequest)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	isEmail := validation.EmailValidation(userLoginRequest.Email)
	if !isEmail {
		utils.ErrorGenerator(w, errorhandling.EmailvalidationError)
		return
	}

	var user response.User
	user, err = a.authService.UserLogin(userLoginRequest)

	if err != nil {
		utils.ErrorGenerator(w, errorhandling.LoginFailedError)
		return
	}

	utils.ResponseGenerator(w, http.StatusOK, user)
}
