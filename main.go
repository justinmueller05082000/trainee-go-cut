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
	charParameter := flag.String("c", "", "Choose characters to print")
	delimitParameter := flag.String("d", "", "Choose a separator")
	flag.Parse()

	if *fieldsParameter != "" && *charParameter != "" {
		fmt.Println("cut: only one type of list may be specified.")
		os.Exit(1)
	} else if *charParameter != "" && *delimitParameter != "" {
		fmt.Println("cut: an input delimiter may be specified only when operating on fields")
		os.Exit(1)
	}

	content, readAllErr := ioutil.ReadAll(os.Stdin)
	if readAllErr != nil {
		fmt.Fprintln(os.Stderr, readAllErr)
		os.Exit(1)
	}

	if len(content) <= 0 {
		return
	}

	var chosenParameter string

	if *charParameter != "" {
		chosenParameter = *charParameter
	} else {
		chosenParameter = *fieldsParameter
	}

	input := strings.Split(chosenParameter, ",")
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
	charCheck := false
	delimiter := []byte{9}

	if *delimitParameter != "" {
		delimiter = []byte(*delimitParameter)
	}

	if *fieldsParameter != "" {
		// Fieldparameter
		contentLineSplit := bytes.Split(content, []byte{'\n'})
		data = make([][][]byte, len(contentLineSplit))

		lineCounter = len(data) - 1

		for i := 0; i < len(contentLineSplit); i++ {
			data[i] = bytes.Split(contentLineSplit[i], delimiter)
		}

		if len(data[lineCounter][0]) > 0 {
			lineCounter = len(data)
		}

		parameterBasedData = data
	} else if *charParameter != "" {
		// Charparameter
		contentCharSplit := bytes.Split(content, []byte("\n"))

		dataCharSplit := make([][][]byte, len(contentCharSplit))

		for i := 0; i < len(contentCharSplit); i++ {
			dataCharSplit[i] = bytes.Split(contentCharSplit[i], []byte(""))
		}

		lineCounter = len(dataCharSplit) - 1

		if len(dataCharSplit[lineCounter]) > 0 {
			lineCounter = len(dataCharSplit)
		}

		charCheck = true
		parameterBasedData = dataCharSplit
	}

	for i := 0; i < lineCounter; i++ {
		lineLength := len(parameterBasedData[i])
		var output []string
		for j := range input {
			for _, k := range inputAsInt[j] {
				if k > lineLength {
					break
				}
				if !charCheck {
					output = append(output, string(parameterBasedData[i][k-1]))
				} else {
					fmt.Printf("%s", parameterBasedData[i][k-1])
				}
			}
		}
		if !charCheck {
			fmt.Println(strings.Join(output, string(delimiter)))
		} else {
			fmt.Printf("\n")
		}
	}
}
