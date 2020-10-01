package config

import (
	"fmt"
	"io/ioutil"
	"log"
)

func GetSettings(dir string, configFileName string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println("Failed to get config file: ", err)
	}

	for _, file := range files {
		if file.Name() == configFileName {
			code, err := ioutil.ReadFile(dir + "/" + configFileName)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(code)
			// parser.Parse(string(code))
		}
	}
}
