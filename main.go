package main

import (
	"strings"
	"fmt"
	"flag"
	"os"
	"bufio"

	printing "github.com/hahwul/gitls/pkg/printing"
	module "github.com/hahwul/gitls/pkg/modules"
)

func main(){
	list := flag.String("l","","List of targets (e.g -l sample.lst)")
	output := flag.String("o","","write output file (optional)")
	version := flag.Bool("version",false,"version of gitls")
	_=output
	flag.Parse()
	if *version {
		fmt.Println(printing.VERSION)
		return
	}

	if *list == "" {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			line := strings.ToLower(sc.Text())
			if line != "" {
				checkURL(line)
			}
		}
	} else {
		target, err := module.ReadLinesOrLiteral(*list)
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
		if strings.Contains(str[2],"github") {
			module.GetRepoListFromUser(str[3], str[2])
		} else if strings.Contains(str[2], "gitlab") {
			// TODO gitlab getting repos
		}
	} else if size == 5 {
		fmt.Println(s)
	}
}
