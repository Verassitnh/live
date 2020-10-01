package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func parse(code string) {
	res := map[string]string{}
	lines := strings.Split(code, "\n")
	for index, line := range lines {

		// Ignore whitespace only lines
		m, err := regexp.MatchString(`^\s*$`, line)
		if err != nil {
			log.Fatal(err)
		}
		if m {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(code, "]") {
			parseKey(line)
		}

		fmt.Println(index, line, res)
	}
}

func parseKey(line string) string {
	var bracketless = strings.TrimSuffix(strings.TrimPrefix(line, "["), "]")
	regex := regexp.MustCompile(`[^a-zA-Z0-9:_\-]`)
	var junkless = regex.ReplaceAllString(line, "-")

	if len(junkless) > 0 {
		return junkless
	} else {
		log.Fatal("Failed to parse config file: ERROR: Empty keys are invalid")
	}
}

func main() {
	parse("dfgn\ndfgdfgdf\ndfgsdgfd\ndfdfn\nsdfsdfsdfsdfsfdsdf\n\n\n\n\n\nsfdsdfgsdfg\nsfdgdsfgds\nshdfgdfdgds\nsdfgsdfgsdfg\nsfsdfgdsfgn\nfdgdsfgsdfg\ndfgdfsgdf\nsdgdfgfn\n  ")
}
