package parser

import (
	"fmt"
	"strings"
)

type Scripts struct {
	defaultScript []string
	scripts       map[string][]string
}

type Settings struct {
	scripts  Scripts
	includes string
	excludes []string
}

func Parse(str string) (settings Settings) {

	keys, values := getValues(str)
	// TODO: Parse keys and values into Settings{}

	fmt.Println(keys)
	fmt.Println(values)

	return Settings{}
}

func getValues(str string) ([]string, [][]string) {
	keys := [][2]int{}
	code := strings.Split(str, "")
	for startIndex, startChar := range code {
		if startChar == "[" {
			for index, char := range code[startIndex:] {
				if char == "]" {
					keys = append(keys, [2]int{startIndex + 1, startIndex + index})
					break
				}
			}
		}
	}
	Keysresult := []string{}
	for _, element := range keys {
		Keysresult = append(Keysresult, strings.Join(code[element[0]:element[1]], ""))
	}
	Valuesresult := [][]string{}

	for index := range keys {
		value := strings.Join(code[keys[index-1][1]:keys[index][0]], "")
		valueArr := strings.Split(value, "\n")

		Valuesresult = append(Valuesresult, valueArr)
	}

	return Keysresult, Valuesresult
}
