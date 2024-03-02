package topics

import (
	"crypto/rand"
	"errors"
	"fmt"
)

func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomBytes := make([]byte, length) //This will give me capital letter Z as the last[8]byte{0, 76, 216, 56, 164, 255, 218, 185}
	result := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println("Error:", err)
		return "", errors.New("could not generate random string")
	}
	for index, value := range randomBytes {
		result[index] = charset[int(value)%len(charset)]
	}
	return string(result), nil
}
