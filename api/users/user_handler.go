package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pranotobudi/go-simple-ecommerce/api"
	"github.com/pranotobudi/go-simple-ecommerce/api/common"
)

type UserRegistrationResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
type UserRegistrationRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserRegistrationResponseFormatter(user User) UserRegistrationResponse {
	formatter := UserRegistrationResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
	return formatter
}

type userHandler struct {
	repository UserRepository
}

func NewUserHandler() *userHandler {
	repository := NewUserRepository()

	return &userHandler{repository}
}

func (h *userHandler) RegisterUser(c echo.Context) error {
	// Input Binding
	userReg := new(UserRegistrationRequest)
	if err := c.Bind(userReg); err != nil {
		return api.ResponseErrorFormatter(c, err)
	}

	// Process Input - User Registration
	user := User{}
	user.Username = userReg.Username
	user.Email = userReg.Email
	user.Password = common.GeneratePassword(userReg.Password)
	savedUser, err := h.repository.AddUser(user)
	if err != nil {
		api.ResponseErrorFormatter(c, err)
	}

	// Success ProductResponse
	data := UserRegistrationResponseFormatter(*savedUser)

	response := api.ResponseFormatter(http.StatusOK, "success", "get user successfull", data)

	return c.JSON(http.StatusOK, response)
}
