package filemanager

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func Encryptor(destinationPath string) Executor {
	key, err := GenerateRandomBytes(32)
	if err != nil {
		log.Println(err)
	}
	log.Println("encryption key: ", base64.URLEncoding.EncodeToString(key))
	return func(sourcefilePath string) {
		file, err := ioutil.ReadFile(sourcefilePath)

		if err != nil {
			log.Println("Readfile error: ", err)
		}

		splitted := strings.Split(sourcefilePath, "/")
		filename := splitted[len(splitted)-1]

		c, err := aes.NewCipher(key)

		if err != nil {
			log.Println(err)
		}

		gcm, err := cipher.NewGCM(c)

		if err != nil {
			fmt.Println(err)
		}

		nonce := make([]byte, gcm.NonceSize())

		if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
			fmt.Println(err)
		}

		encryptedFile := gcm.Seal(nonce, nonce, file, nil)

		err = ioutil.WriteFile(destinationPath+"/"+filename+".encrypted", encryptedFile, 0644)
		if err != nil {
			log.Println("Readfile error: ", err)
		}
		log.Println("Encryption is successfull")

	}
}

// DecryptFile decrypts an encrypted file
func DecryptFile(filePath string, keyStr string) ([]byte, error) {
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
