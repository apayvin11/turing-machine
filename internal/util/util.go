package util

import (
	"bufio"
	"os"
)

// ReadFileIntoSliceByLines reads the file line by line into a slice
func ReadFileIntoSliceByLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := []string{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		res = append(res, sc.Text())
	}
	return res, nil
}
