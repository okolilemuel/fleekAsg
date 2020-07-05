package encryptor

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
