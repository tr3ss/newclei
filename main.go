package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

func getPRs(token string) []float64 {

	// Get request
	var result []float64
	page := 1
	for {
		ghprs := "https://api.github.com/repos/projectdiscovery/nuclei-templates/pulls?state=open&per_page=100&page=" + strconv.Itoa(page)
		req, _ := http.NewRequest("GET", ghprs, nil)
		req.Header.Set("Host", "api.github.com")
		if token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		json := string(body[:])
		if strings.TrimRight(json, "\n") == "[]" {
			break
		}
		prs := gjson.Get(json, "#(title)#.number").Array()
		for _, v := range prs {
			result = append(result, v.Num)
		}
		page++
	}
	return result
}

func getFiles(prs float64, token string, cves bool) gjson.Result {

	// Get request
	prstr := fmt.Sprintf("%v", prs)

	ghfiles := "https://api.github.com/repos/projectdiscovery/nuclei-templates/pulls/" + prstr + "/files"
	req, _ := http.NewRequest("GET", ghfiles, nil)
	req.Header.Set("Host", "api.github.com")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json := string(body[:])
	var files gjson.Result
	if cves {
		files = gjson.Get(json, "#(filename%\"*cves*\")#.raw_url")
	} else {
		files = gjson.Get(json, "#(filename)#.raw_url")
	}
	return files
}

func main() {

	var token string
	flag.StringVar(&token, "token", "", "- GitHub token to use.\n")
	var cves bool
	flag.BoolVar(&cves, "cves", false, "- Show only CVEs")
	flag.Parse()

	prs := getPRs(token)
	var files gjson.Result
	for i := range prs {
		pr := prs[i]
		files = getFiles(pr, token, cves)
		for _, name := range files.Array() {
			println(name.String())
		}
	}
}
