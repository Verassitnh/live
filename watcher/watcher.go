package watcher

import (
	"io/ioutil"
	"log"

	"github.com/fsnotify/fsnotify"
)

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
		err = watcher.Add(dir)
		if err != nil {
			log.Fatal(err)
		}
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

	for _, file := range files {

		if file.IsDir() {
			err := watcher.Add(dir + "/" + file.Name())
			if err != nil {
				log.Printf("FAILED ADDING FOLDER: \"%s/%s\"! ERROR: %s", dir, file.Name(), err)
			}
			recAddFiles(dir+"/"+file.Name(), watcher)
		}
	}

}
