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
	flag.Parse()

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

	buf := bufio.NewReader(os.Stdin)
	var data [][]byte
	delimiter := []byte{9}

	for {
		content, readAllErr := buf.ReadBytes('\n')
		if readAllErr != nil && readAllErr != io.EOF {
			fmt.Fprintln(os.Stderr, readAllErr)
			os.Exit(1)
		}

		contentLineSplit := bytes.Split(content, []byte{'\n'})
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
