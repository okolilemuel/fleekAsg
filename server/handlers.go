package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/okolilemuel/fleekAsg/filemanager"
)

type httpHandler func(w http.ResponseWriter, r *http.Request)

func getFile(path string) httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Query().Get("filename")
		key := r.URL.Query().Get("key")
		if filename != "" && key != "" {
			log.Println(filename, key)
		}

		file, err := filemanager.DecryptFile(path+"/"+filename, key)
		if err != nil {
			log.Println(err)
			respondWithJSON(nil, http.StatusOK, false, "File decryption error", w)
			return
		}
		fileNameParts := strings.Split(filename, ".")
		downloadFileName := fileNameParts[0] + "." + fileNameParts[1]
		log.Println(downloadFileName)
		if err != nil {
			log.Println("Readfile error: ", err)
		}
		respondWithFile(w, file, downloadFileName)
	}
}

func getFiles(path string) httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		files, err := ioutil.ReadDir(path + "/")
		if err != nil {
			log.Println("Error occured while reading directory", err)
		}

		fileNames := []string{}
		for _, f := range files {
			fileNames = append(fileNames, f.Name())
		}
		data := struct{ FileNames []string }{
			FileNames: fileNames,
		}
		respondWithJSON(data, http.StatusOK, true, "Get files route", w)
	}
}
