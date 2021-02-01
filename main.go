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
		module.GetRepoListFromUser(str[3])
	} else if size == 5 {
		fmt.Println(s)
	}
}
