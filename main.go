package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fieldsParameter := flag.String("f", "", "Choose field to print")
	charParameter := flag.String("c", "", "Choose characters to print")
	delimitParameter := flag.String("d", "", "Choose a separator")
	onlyDelimitParameter := flag.Bool("s", false, "Suppress lines with no field delimiter characters. Unless specified, lines with no delimiters are passed through unmodified.")
	flag.Parse()

	if (*fieldsParameter != "") == (*charParameter != "") {
		fmt.Fprintln(os.Stderr, "cut: only one type of list may be specified.")
		os.Exit(1)
	} else if *charParameter != "" && *delimitParameter != "" {
		fmt.Fprintln(os.Stderr, "cut: an input delimiter may be specified only when operating on fields")
		os.Exit(1)
	} else if *charParameter != "" && *onlyDelimitParameter {
		fmt.Fprintln(os.Stderr, "cut: suppressing non-delimited lines makes sense\n\tonly when operating on fields")
		os.Exit(1)
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

	buf := bufio.NewReader(os.Stdin)
	var data [][]byte
	delimiter := []byte{9}

	if *delimitParameter != "" {
		delimiter = []byte(*delimitParameter)
	} else if *charParameter != "" {
		delimiter = []byte{}
	}

	for {
		content, readAllErr := buf.ReadBytes('\n')
		if readAllErr != nil && readAllErr != io.EOF {
			fmt.Fprintln(os.Stderr, readAllErr)
			os.Exit(1)
		}

		contentLineSplit := bytes.Split(content, []byte{'\n'})

		if !*onlyDelimitParameter || bytes.Contains(contentLineSplit[0], delimiter) {
			data = bytes.Split(contentLineSplit[0], delimiter)

			lineLength := len(data)
			var output [][]byte

			for j := range input {
				for _, k := range inputAsInt[j] {
					if k > lineLength {
						break
					}
					output = append(output, data[k-1])
				}
			}

			contentLineSplit[0] = bytes.Join(output, delimiter)

			_, _ = os.Stdout.Write(bytes.Join(contentLineSplit, []byte{'\n'}))

			if readAllErr == io.EOF {
				break
			}
		}
	}
}

