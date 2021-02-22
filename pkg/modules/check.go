package modules

import (
	"fmt"
	"strings"

	"github.com/hahwul/gitls/pkg/model"
)

// CheckAccount is repo list of accounts in target
func CheckAccount(s string, options model.Options) {
	str := strings.Split(s, "/")
	size := len(str) // 4 is user/org , 5 is repository
	if size == 4 {
		if strings.Contains(str[2], "github") {
			GetRepoListFromIncludeAccount(str[3], str[2], options)
		} else if strings.Contains(str[2], "gitlab") {
			// TODO gitlab getting repos
		}
	} else if size == 5 {
		fmt.Println(s)
	}
}

// CheckURL is repo list of target
func CheckURL(s string, options model.Options) {
	str := strings.Split(s, "/")
	size := len(str) // 4 is user/org , 5 is repository
	if size == 4 {
		if strings.Contains(str[2], "github") {
			GetRepoListFromUser(str[3], str[2], options)
		} else if strings.Contains(str[2], "gitlab") {
			// TODO gitlab getting repos
		}
	} else if size == 5 {
		fmt.Println(s)
	}
}
