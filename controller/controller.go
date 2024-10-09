package controller

import (
	"auth-go-app/db"
	"auth-go-app/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func RegisterUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateModel(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": formatValidationErrors(err)})
		return
	}

	if db.CheckIfUserExists(user.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "User with this email already exists"})
		return
	}

	requestBody, err := createRequestBody(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create request body for auth: " + err.Error()})
		return
	}

	body, err := makeRequest("POST", os.Getenv("AUTH_URL"), requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var data models.Response
	if err := json.Unmarshal(body, &data); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse Auth0 response"})
		return
	}

	if data.StatusCode > 0 {
		ctx.JSON(data.StatusCode, gin.H{"message": "Error from Auth0: " + data.Message})
		return
	}

	if strings.EqualFold(data.Email, user.Email) {
		user.ID = data.ID
		user.Password = ""
		if err := db.Save(&user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save user: " + err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "User has been registered"})
	}
}

func createRequestBody(user models.User) ([]byte, error) {
	jsonData := map[string]interface{}{
		"email":      user.Email,
		"password":   user.Password,
		"connection": os.Getenv("AUTH_DB"),
	}
	return json.Marshal(jsonData)
}

func makeRequest(method, url string, jsonBody []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func validateModel(user interface{}) error {
	return validate.Struct(user)
}

func formatValidationErrors(err error) string {
	var errParts []string
	for _, err := range err.(validator.ValidationErrors) {
		errParts = append(errParts, fmt.Sprintf("%s %s", err.Field(), err.Tag()))
	}
	return "Validation errors: " + strings.Join(errParts, ", ")
}
