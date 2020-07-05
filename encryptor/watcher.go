package encryptor

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

type Executor func(filelocation string)

func Watcher(sourcePath string, executor Executor) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					executor(event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(sourcePath)
	if err != nil {
		return err
	}
	<-done
	return nil
}
