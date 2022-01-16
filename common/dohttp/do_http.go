package dohttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var regxTitle = regexp.MustCompile(`<title>([^<]+)</title>`)
var regxKWords = regexp.MustCompile(`<meta name="keywords" content="([^"]+)" />`)
var regxDesc = regexp.MustCompile(`<meta name="description" content="([^"]+)" />`)
var regxUrl = regexp.MustCompile(`(https?://[^/^"^\?]+)`)

func Do(url string, urlMap map[string]bool) (title, kWords, desc string) {
	cli := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	resp, err := cli.Do(req)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	bodyStr := string(body)

	reRst := regxTitle.FindAllStringSubmatch(bodyStr, 1)
	if len(reRst) > 0 {
		title = reRst[0][1]
	}
	reRst = regxKWords.FindAllStringSubmatch(bodyStr, 1)
	if len(reRst) > 0 {
		kWords = reRst[0][1]
	}
	reRst = regxDesc.FindAllStringSubmatch(bodyStr, 1)
	if len(reRst) > 0 {
		desc = reRst[0][1]
	}
	//fmt.Println(bodyStr)
	fmt.Printf("##########\nU:%s\nT:%s\nK:%s\nD:%s\n", url, title, kWords, desc)

	reRst = regxUrl.FindAllStringSubmatch(bodyStr, -1)
	for _, reIt := range reRst {
		newURL := reIt[1]
		_, ok := urlMap[newURL]
		if ok {
			continue
		}
		urlMap[newURL] = false
	}
	fmt.Println(len(urlMap))

	return
}
