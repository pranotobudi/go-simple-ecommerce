package users

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/pranotobudi/go-simple-ecommerce/api"
	"github.com/pranotobudi/go-simple-ecommerce/common"

	"golang.org/x/crypto/bcrypt"
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

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	AuthToken string `json:"auth_token"`
}

func UserResponseFormatter(user User, auth_token string) UserLoginResponse {
	formatter := UserLoginResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		AuthToken: auth_token,
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

func (h *userHandler) UserLogin(c echo.Context) error {
	// Input Binding
	userLogin := UserLoginRequest{}
	if err := c.Bind(&userLogin); err != nil {
		return api.ResponseErrorFormatter(c, err)
	}

	// Process Input
	authUser, err := h.AuthUser(userLogin)
	fmt.Println("We're IN HERE: USERLOGIN INSIDE: authUser: ", authUser)
	if err != nil {
		return api.ResponseErrorFormatter(c, err)
	}

	// Create JWT token
	auth_token, err := h.CreateAccessToken()
	if err != nil {
		return api.ResponseErrorFormatter(c, err)
	}

	// Success UserLoginResponse
	data := UserResponseFormatter(*authUser, auth_token)
	response := api.ResponseFormatter(http.StatusOK, "success", "user login successfull", data)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) AuthUser(req UserLoginRequest) (*User, error) {
	username := req.Username
	password := req.Password
	fmt.Println("AUTHUSER CALLED, username: ", username, " password: ", password)

	//check Author Table
	user, err := h.repository.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("username is not registered")
	}

	test, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Printf("COMPARES: %s %s \n", user.Password, string(test))
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}
	return user, nil
}

func (s *userHandler) CreateAccessToken() (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedKey, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return signedKey, err
	}

	return signedKey, nil
}
