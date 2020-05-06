package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fieldsParameter := flag.String("f", "", "Choose field to print")
	flag.Parse()

	content, readAllErr := ioutil.ReadAll(os.Stdin)
	if readAllErr != nil {
		fmt.Fprintln(os.Stderr, readAllErr)
		os.Exit(1)
	}

	if len(content) <= 0 {
		return
	}

	input := strings.Split(*fieldsParameter, ",")
	inputAsInt := make([][]int, len(input))

	for i := range input {
		split := strings.Split(input[i], "-")
		inputAsInt[i] = make([]int, len(split))

		for sid, s := range split {
			value, err := strconv.Atoi(s)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			inputAsInt[i][sid] = value
		}

		if len(inputAsInt[i]) > 1 && inputAsInt[i][1]-inputAsInt[i][0] > 1 {
			for p := inputAsInt[i][0] + 1; p < inputAsInt[i][1]; p++ {
				inputAsInt[i] = append(inputAsInt[i], p)
			}
			sort.Ints(inputAsInt[i])
		}
	}

	var data, parameterBasedData [][][]byte
	var lineCounter int
	delimiter := []byte{9}

	// Fieldparameter
	contentLineSplit := bytes.Split(content, []byte{'\n'})
	data = make([][][]byte, len(contentLineSplit))

	for i := 0; i < len(contentLineSplit); i++ {
		data[i] = bytes.Split(contentLineSplit[i], delimiter)
	}

	lineCounter = len(data) - 1

	if len(data[lineCounter][0]) > 0 {
		lineCounter = len(data)
	}

	parameterBasedData = data

	for i := 0; i < lineCounter; i++ {
		lineLength := len(parameterBasedData[i])
		var output []string
		for j := range input {
			for _, k := range inputAsInt[j] {
				if k > lineLength {
					break
				}
				output = append(output, string(parameterBasedData[i][k-1]))
			}
		}
		fmt.Println(strings.Join(output, string(delimiter)))
	}
}
