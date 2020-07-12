# Fleek Asg

This app takes a `source path` and `destination path` at run time. It encrypts any file placed in the source path and stores it in the destination path. The encryption key for this particular session will be logged to the terminal. The app exposes http endpoints to query the encrypted files and download the decrypted version.

# Usage

Follow the steps below to build and run the app

- Clone this repository 
- Open the terminal or command line in the base directory of the project
- Build the binary using `go build -ldflags="-w -s" -o ./bin/fleekAsg ./main/main.go`
- Run the application using `./bin/fleekAsg <sourcePath> <destinationPath>`
- To get the list of encrypted files call `GET http://localhost:8080/files/`
- To download a decrypted file, call `http://localhost:8080/file/?key=<encryptionKey>&filename=<fileNme.encrypted>`
    - eg `http://localhost:8080/file/?key=T8dL_xp6WvySO-Duthht6a40v5LrTFfHtAQHbQjJ3xs=&filename=file1.encrypted`

Note: the encryption key is displayed on the terminal after the application has started running

# Dependencies

- [fsnotify](github.com/fsnotify/fsnotify) 

# Credits

- https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
- https://blog.questionable.services/article/generating-secure-random-numbers-crypto-rand/
- https://github.com/fsnotify/fsnotify
- https://stackoverflow.com/questions/24116147/how-to-download-file-in-browser-from-go-server
- https://stackoverflow.com/questions/14668850/list-directory-in-go
