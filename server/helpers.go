package server

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

func respondWithJSON(data interface{}, statusCode int, success bool, message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	httpResponse := response{
		Success: success,
		Data:    data,
		Message: message,
	}

	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(httpResponse)

	if err != nil {
		log.Println("[ERROR] respondWithJSON() : encountered and error while converting data to JSON")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}

}

func respondWithFile(w http.ResponseWriter, file []byte, filename string) {
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))
	w.Header().Set("Content-Type", "application/octet-stream")

	w.Write(file)
}

func decryptFile(filePath string, keyStr string) ([]byte, error) {
	encryptedFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	key, _ := base64.URLEncoding.DecodeString(keyStr)
	// key := []byte(keyStr)
	c, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedFile) < nonceSize {
		return nil, err
	}

	nonce, ciphertext := encryptedFile[:nonceSize], encryptedFile[nonceSize:]
	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)

	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
