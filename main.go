package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"io"
)

type CSVReader csv.Reader

func countLines(r io.Reader) int {
	count := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		count++
	}
	return count
}

func ReadAll(r io.Reader) ([][]string, error) {
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)
	size := countLines(tee)
	csvReader := csv.NewReader(&buf)
	records := make([][]string, 0, size)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}
