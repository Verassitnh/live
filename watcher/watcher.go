package watcher

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fsnotify/fsnotify"
)

// type FSWatcherEvent struct {
// 	RelativePath string
// 	Path         string
// 	FileName     string
// 	EventName    string
// 	Time         string
// }

func Watch(dir string, recursive bool, channel chan fsnotify.Event) {

	// Start Watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
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
				channel <- event

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Fatal(err)

			}

		}
	}()

	// Add files
	if recursive == true {
		recAddFiles(dir, watcher)
	} else {
		err := watcher.Add(dir)
		if err != nil {
			log.Fatal(err)
		}
	}

	// This says its done watching??
	<-done
}

func recAddFiles(dir string, watcher *fsnotify.Watcher) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if file.IsDir() {
			recAddFiles(dir+"/"+file.Name(), watcher)
			err := watcher.Add(file.Name())
			if err != nil {
				log.Fatalf("FAILED ADDING FOLDER: \"%s/%s\" WITH: %s", dir, file.Name(), err)
			}
			fmt.Println("Added Folder:", file.Name())
		}
		fmt.Println("Skipping File:", file.Name())

	}

}
