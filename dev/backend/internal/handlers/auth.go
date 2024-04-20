package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/skapskap/iheros_tavern/internal/data"
	"github.com/skapskap/iheros_tavern/utility"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func AuthRegister(c echo.Context, db data.DBTX) error {
	var input struct {
		Username string
		Email    string
		Password string
	}

	if err := c.Bind(&input); err != nil {
		return utility.JSONResponse(c, http.StatusBadRequest, "Error", err.Error(), nil)
	}

	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		return utility.JSONResponse(c, http.StatusInternalServerError, "Error", "Failed to encrypt password", nil)
	}

	var params = data.CreateUserParams{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
	}

	user, err := data.New(db).CreateUser(c.Request().Context(), params)
	if err != nil {
		return utility.JSONResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
	}

	return utility.JSONResponse(c, http.StatusCreated, "Created", "User created successfully", user)
}

func AuthLogin(c echo.Context, db data.DBTX) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&input); err != nil {
		return utility.JSONResponse(c, http.StatusBadRequest, "Error", "Invalid request data", nil)
	}

	user, err := data.New(db).GetUserByEmail(c.Request().Context(), input.Email)
	if err != nil {
		return utility.JSONResponse(c, http.StatusInternalServerError, "Error", "User not found", nil)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return utility.JSONResponse(c, http.StatusUnauthorized, "Error", "Invalid credentials", nil)
	}

	token, err := utility.GenerateJWT(user.ID, user.Username)
	if err != nil {
		return utility.JSONResponse(c, http.StatusInternalServerError, "Error", "Failed to generate token", nil)
	}

	userResponse := struct {
		ID        int32  `json:"id"`
		Email     string `json:"email"`
		Username  string `json:"username"`
	}{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
	}

	response := struct {
		User  interface{} `json:"user"`
		Token string      `json:"token"`
	}{
		User:  userResponse,
		Token: token,
	}

	return utility.JSONResponse(c, http.StatusOK, "Success", "Logged in successfully", response)
}
