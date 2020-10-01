package main

import (
	"fmt"
	"log"
	"os"

	sett "github.com/Verassitnh/live/config"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	sett.GetSettings(dir, ".liverc")
	fmt.Println("Hello World!")

	// NOTE: This is commented out to test sett.GetSettings()

	// fmt.Print("What command would you like to run?\nEnter text: ")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// command := make([]string, 0)

	// for _, element := range strings.Split(scanner.Text(), " ") {
	// 	command = append(command, element)
	// }

	// go exec.Command(command)

	// chann := make(chan fsnotify.Event)

	// go fswatch.Watch(dir, true, chann)
	// fmt.Println("Watching")

	// for range chann {
	// 	log.Printf("\n\n####:: restarted process %q ::####\n\n", command[0])
	// 	exec.Command(command)
	// }
}
