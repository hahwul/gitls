package main

import (
	"strings"
	"flag"
	"os"
	"bufio"

	// "github.com/hahwul/gitls/modules"
)

func main(){
	list := flag.String("l","","List of targets (e.g Repository URL and Owner URL and User URL)")
	output := flag.String("o","","write output file (optional)")
	var targets []string
	_=targets
	_=output

	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}
	if *list == "" {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			line := strings.ToLower(sc.Text())
			if line != "" {
				targets = append(targets, line)
			}
		}
	} else {
		flag.Usage()
		return
	}
	// m.GetRepoListFromUser("hahwul")
}
