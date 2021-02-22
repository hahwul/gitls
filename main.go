package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	model "github.com/hahwul/gitls/pkg/model"
	module "github.com/hahwul/gitls/pkg/modules"
	printing "github.com/hahwul/gitls/pkg/printing"
)

func main() {
	list := flag.String("l", "", "List of targets (e.g -l sample.lst)")
	output := flag.String("o", "", "write output file (optional)")
	version := flag.Bool("version", false, "version of gitls")
	proxy := flag.String("proxy", "", "using custom proxy")
	useTor := flag.Bool("tor", false, "using tor proxy / localhost:9050")
	includeAccount := flag.Bool("include-users", false, "include repo of org users(member)")
	flag.Parse()
	options := model.Options{
		Proxy:          *proxy,
		UseTor:         *useTor,
		Output:         *output,
		IncludeAccount: *includeAccount,
	}
	if *version {
		fmt.Println(printing.VERSION)
		return
	}

	if *list == "" {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			line := strings.ToLower(sc.Text())
			if line != "" {
				module.CheckURL(line, options)
			}
			if *includeAccount {
				module.CheckAccount(line, options)
			}
		}
	} else {
		target, err := module.ReadLinesOrLiteral(*list)
		_ = err
		for i, v := range target {
			_ = i
			module.CheckURL(v, options)
			if *includeAccount {
				module.CheckAccount(v, options)
			}
		}

	}
}
