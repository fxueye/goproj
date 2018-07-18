package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

func MakeMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	s := hex.EncodeToString(h.Sum(nil))
	return s
}

func MakeGetParams(params map[string]interface{}) (string, error) {
	keys := make([]string, len(params))
	i := 0
	for k, _ := range params {
		keys[i] = k
		i++
	}
	retStr := ""
	sort.Strings(keys)
	for j := 0; j < len(keys); j++ {
		key := keys[j]
		if j == 0 {
			retStr += "" + key + "=" + params[key].(string)
		} else {
			retStr += "&" + key + "=" + params[key].(string)
		}
	}
	return retStr, nil
}
func HttpGet(baseUrl string, params map[string]interface{}) (string, error) {
	paramsStr, err := MakeGetParams(params)
	url := fmt.Sprintf("%s?%s", baseUrl, paramsStr)

	body, err := http.Get(url)
	defer body.Body.Close()
	respBody, err := ioutil.ReadAll(body.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), err
}

func HttpPostAndHeader(baseUrl string, params map[string]interface{}, headers map[string]interface{}) (string, error) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    5 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	var r http.Request
	r.ParseForm()
	for k, v := range params {
		r.Form.Add(k, v.(string))
	}
	bodystr := strings.TrimSpace(r.Form.Encode())
	req, err := http.NewRequest("POST", baseUrl, strings.NewReader(bodystr))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Connection", "Keep-Alive")
	for k, v := range headers {
		req.Header.Set(k, v.(string))
	}
	body, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer body.Body.Close()
	respBody, err := ioutil.ReadAll(body.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), err

}

func HttpPost(baseUrl string, params map[string]interface{}) (string, error) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    5 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	var r http.Request
	r.ParseForm()
	for k, v := range params {
		r.Form.Add(k, v.(string))
	}
	bodystr := strings.TrimSpace(r.Form.Encode())
	req, err := http.NewRequest("POST", baseUrl, strings.NewReader(bodystr))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Connection", "Keep-Alive")
	body, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer body.Body.Close()
	respBody, err := ioutil.ReadAll(body.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), err

}
func GetStrSign(data map[string]interface{}) (string, error) {
	keys := make([]string, len(data))
	i := 0
	for k, _ := range data {
		keys[i] = k
		i++
	}
	retStr := ""
	sort.Strings(keys)
	for j := 0; j < len(keys); j++ {
		key := keys[j]
		if j == 0 {
			retStr += "" + key + "=" + data[key].(string)
		} else {
			retStr += "&" + key + "=" + data[key].(string)
		}
	}
	return retStr, nil
}
