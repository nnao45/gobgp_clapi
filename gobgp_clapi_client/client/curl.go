package client

import (
	"bytes"
	"fmt"
	"github.com/mattn/go-scan"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func curl_check_jwt(auth1 string, auth2 string) bool {
        req, err := http.NewRequest("GET", GOBGP_JWTSTATUS, nil)
        if err != nil {
                return false
        }
//        req.SetBasicAuth(auth1, auth2)
        req.Header.Set(auth1, auth2)

        resp, err := http.DefaultClient.Do(req)
        if err != nil {
                return false
        }
        if resp.StatusCode == 200 {
                return true
        }
	defer resp.Body.Close()
        return false
}


func curl_check(auth1 string, auth2 string) bool {
	req, err := http.NewRequest("GET", GOBGP_STATUS, nil)
	if err != nil {
		return false
	}
	req.SetBasicAuth(auth1, auth2)
//	req.Header.Set(auth1, auth2)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}
	if resp.StatusCode == 200 {
		return true
	}
	defer resp.Body.Close()
	return false
}

func curl_get(values url.Values, auth1 string, auth2 string) string {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	req, err := http.NewRequest("GET", GOBGP_TOKEN, nil)
	if err != nil {
		return ""
	}
	req.SetBasicAuth(auth1, auth2)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	var js = strings.NewReader(execute(resp))
	var s string

	if err := scan.ScanJSON(js, "/token/", &s); err != nil {
		return ""
	}
	return s

}

func curl_post_command(values url.Values, hash string) {

	jsondata := bytes.NewBuffer([]byte(cat(LCOMMANDFILE)))
	 req, err := http.NewRequest("POST", GOBGP_COMMAND, jsondata)
	if err != nil {
		fmt.Println(err)
	}
//	req.SetBasicAuth(hash, "unused")
	req.Header.Set("Content-Type", "application/json")
	hash = "Bearer " + hash
        req.Header.Set("Authorization", hash)
//	 req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)

	}
	defer resp.Body.Close()
	execute(resp)

}

func execute(resp *http.Response) string {
	// response bodyを文字列で取得するサンプル
	// ioutil.ReadAllを使う
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		return string(b)
	}
	return ""
}
