package handlers

import (
	"errors"
	"fmt"
	"net/http"
	models "moshenskyi/caesar/pkg/models"

	"github.com/gin-gonic/gin"
)

func encryptInternal(message string, shift uint8) (string, error) {
	runes := []rune(message)

	for index, char := range message {
		if char >= 'a' && char <= 'z' {
			runes[index] = (char-'a'+rune(shift))%26 + 'a'
		} else if char >= 'A' && char <= 'Z' {
			runes[index] = (char-'A'+rune(shift))%26 + 'A'
		} else {
			return "", errors.New("Non-ASCII character detected: " + string(char))
		}

	}

	return string(runes), nil
}

func Encrypt(context *gin.Context) {
	var message models.EncryptionData

	if err := context.BindJSON(&message); err != nil {
		return
	}

	result, err := encryptInternal(message.Value, uint8(message.Shift))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	context.IndentedJSON(http.StatusOK, result)
}
