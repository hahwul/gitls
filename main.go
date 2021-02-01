package main

import (
	"strings"
	"fmt"
	"flag"
	"os"
	"bufio"

	module "github.com/hahwul/gitls/pkg/modules"
)

func main(){
	list := flag.String("l","","List of targets (e.g Repository URL and Owner URL and User URL)")
	output := flag.String("o","","write output file (optional)")
	_=output

	flag.Parse()
	if *list == "" {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			line := strings.ToLower(sc.Text())
			if line != "" {
				checkURL(line)
			}
		}
	} else {
		target, err := readLinesOrLiteral(*list)
		_ = err
		for i, v := range target {
			_ = i
			checkURL(v)
		}
	}
}
func checkURL(s string) {
	str := strings.Split(s,"/")
	size := len(str) // 4 is user/org , 5 is repository
	if size == 4 {
		module.GetRepoListFromUser(str[3])
	} else if size == 5 {
		fmt.Println(s)
	}
}

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

// readLinesOrLiteral tries to read lines from a file, returning
// the arg in a string slice if the file doesn't exist, unless
// the arg matches its default value
func readLinesOrLiteral(arg string) ([]string, error) {
	if isFile(arg) {
		return readLines(arg)
	}

	// if the argument isn't a file, but it is the default, don't
	// treat it as a literal value

	return []string{arg}, nil
}

// isFile returns true if its argument is a regular file
func isFile(path string) bool {
	f, err := os.Stat(path)
	return err == nil && f.Mode().IsRegular()
}
