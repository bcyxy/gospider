package main

import (
	"fmt"

	"github.com/bcyxy/gospider/common/dohttp"
	"github.com/bcyxy/gospider/common/glbval"
)

func main() {
	fmt.Println("Hello gospider")

	fmt.Println(glbval.GitCommitID)
	fmt.Println(glbval.BuildTime)

	urlMap := make(map[string]bool)
	//url := "https://ask.zol.com.cn" //乱码
	url := "http://www.techweb.com.cn"
	dohttp.Do(url, urlMap)
	for {
		urlList := []string{}
		for url, done := range urlMap {
			if done {
				continue
			}
			urlList = append(urlList, url)
		}
		for _, url := range urlList {
			urlMap[url] = true
			dohttp.Do(url, urlMap)
		}
	}
}
