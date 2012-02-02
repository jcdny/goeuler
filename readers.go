package euler

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func TableIntRead(in *bufio.Reader) (table [][]int, err error) {
	lines := 0
OUT:
	for {
		var line string

		line, err = in.ReadString('\n')
		lines++
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		words := strings.Split(line, " ")
		row := make([]int, 0, len(words))
		for _, word := range words {
			val, err := strconv.Atoi(word)
			if err != nil {
				log.Printf("Conversion error line %v : %v", lines, err)
				break OUT
			}
			row = append(row, val)
		}
		table = append(table, row)
	}

	return
}

func TableIntReadFile(file string) (table [][]int, err error) {
	var f *os.File
	f, err = os.Open(file)
	if err != nil {
		return
	} else {
		defer f.Close()
		in := bufio.NewReader(f)
		table, err = TableIntRead(in)
		if err == io.EOF {
			err = nil
		}
	}

	return
}
