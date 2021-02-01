package main

import (
	"strings"
	"fmt"
	"flag"
	"os"
	"bufio"

	printing "github.com/hahwul/gitls/pkg/printing"
	module "github.com/hahwul/gitls/pkg/modules"
	model "github.com/hahwul/gitls/pkg/model"
)

func main(){
	list := flag.String("l","","List of targets (e.g -l sample.lst)")
	output := flag.String("o","","write output file (optional)")
	version := flag.Bool("version",false,"version of gitls")
	proxy := flag.String("proxy","","using custom proxy")
	useTor := flag.Bool("tor",false,"using tor proxy / localhost:9050")
	options := model.Options{
		Proxy:            *proxy,
		UseTor:           *useTor,
		Output:           *output,
	}
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
				checkURL(line, options)
			}
		}
	} else {
		target, err := module.ReadLinesOrLiteral(*list)
		_ = err
		for i, v := range target {
			_ = i
			checkURL(v, options)
		}
	}
}
func checkURL(s string, options model.Options) {
	str := strings.Split(s,"/")
	size := len(str) // 4 is user/org , 5 is repository
	if size == 4 {
		if strings.Contains(str[2],"github") {
			module.GetRepoListFromUser(str[3], str[2], options)
		} else if strings.Contains(str[2], "gitlab") {
			// TODO gitlab getting repos
		}
	} else if size == 5 {
		fmt.Println(s)
	}
}
