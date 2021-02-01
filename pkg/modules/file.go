package modules

import (
	"os"
	"bufio"
)

func readLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()

	lines := make([]string, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines, sc.Err()
}

// ReadLinesOrLiteral tries to read lines from a file, returning
func ReadLinesOrLiteral(arg string) ([]string, error) {
	if isFile(arg) {
		return readLines(arg)
	}
	return []string{arg}, nil
}

// isFile returns true if its argument is a regular file
func isFile(path string) bool {
	f, err := os.Stat(path)
	return err == nil && f.Mode().IsRegular()
}
