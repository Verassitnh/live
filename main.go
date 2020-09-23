package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	fswatch "github.com/Verassitnh/live/watcher"
	"github.com/fsnotify/fsnotify"
)

func main() {
	fmt.Print("What command would you like to run?\nEnter text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	command := make([]string, 0)

	for _, element := range strings.Split(scanner.Text(), " ") {
		command = append(command, element)
	}

	runCommand(command)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	chann := make(chan fsnotify.Event)

	go fswatch.Watch(dir, true, chann)

	for range chann {
		log.Printf("\n\n####:: restarted proccess %q ::####\n\n", command[0])
		runCommand(command)
	}
}

func runCommand(command []string) {
	var cmd *exec.Cmd
	if len(command) > 1 {
		cmd = exec.Command(command[0], command[1:]...)
	} else {
		cmd = exec.Command(command[0])
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
