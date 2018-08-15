package main

import (
	"flag"
	"fmt"
	"game/common/wx"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var (
	http  *wx.Client
	mu    sync.RWMutex
	pages map[string]interface{}
)

func main() {
	web_url := flag.String("U", "https://www.baidu.com", "web base url ")
	out_path := flag.String("O", "./web", "web html out path")
	flag.Parse()
	http = wx.NewClient()
	pages = make(map[string]interface{})
	webUrl := fmt.Sprintf("%s", *web_url)
	fmt.Printf("web url:%v\n", webUrl)
	outPath, _ := filepath.Abs(*out_path)
	fmt.Printf("out path:%v\n", outPath)
	os.RemoveAll(outPath)
	if !IsExist(outPath) {
		os.MkdirAll(outPath, 0755)
	}
	getPageAllData(outPath, webUrl, webUrl)

	chSig := make(chan os.Signal)
	signal.Notify(chSig)

	for {
		select {
		case sig := <-chSig:
			{
				fmt.Printf("over:%v", sig)
			}
		}
	}
}
func getPageAllData(outPath, baseUrl, webUrl string) {
	fmt.Printf("outPath:%v\n", outPath)
	fmt.Printf("baseUrl:%v\n", baseUrl)
	fmt.Printf("webUrl:%v\n", webUrl)
	mu.RLock()
	defer mu.RUnlock()

	paths, filename := filepath.Split(webUrl)

	filePath := ""
	if paths == "http://" || paths == "https://" {
		filePath = fmt.Sprintf("%s/index.html", outPath)
	} else {
		index := strings.Index(filename, "?")
		if index == 0 {
			re := regexp.MustCompile(`\?mod=(\w+)&`)
			finds := re.FindAllStringSubmatch(webUrl, -1)
			if len(finds) > 0 {
				filename = finds[0][1]
			}
		}
		filePath = fmt.Sprintf("%s/%s.html", outPath, filename)

	}
	if IsExist(filePath) {
		return
	}
	ret, err := http.Get(webUrl, nil)
	if err != nil {
		fmt.Printf("http error %v", err)
		return
	}
	content := string(ret)
	createFile(filePath, content)
	allJsPath := getAllJavascript(content)
	for _, path := range allJsPath {
		urlPath := path
		oPath := path
		index := strings.Index(path, "http")
		if index == -1 {
			urlPath = fmt.Sprintf("%s%s", baseUrl, path)
			oPath = fmt.Sprintf("%s%s", outPath, path)
		} else {
			oPath = fmt.Sprintf("%s/js/%s", outPath, filepath.Base(path))
		}
		ret, err = http.Get(urlPath, nil)
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}
		content := string(ret)
		fmt.Printf("js####:%v\n", urlPath)
		createFile(oPath, content)
	}
	allCssPath := getAllCss(content)
	for _, path := range allCssPath {
		urlPath := path
		oPath := path
		index := strings.Index(path, "http")
		if index == -1 {
			urlPath = fmt.Sprintf("%s/%s", baseUrl, path)
			oPath = fmt.Sprintf("%s%s", outPath, path)
		} else {
			oPath = fmt.Sprintf("%s/css/%s", outPath, filepath.Base(path))
		}
		ret, err = http.Get(urlPath, nil)
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}
		content := string(ret)
		createFile(oPath, content)
	}
	allImg := getAllImg(content)
	fmt.Printf("allImg:%v\n", allImg)
	for _, path := range allImg {
		urlPath := path
		oPath := path
		index := strings.Index(path, "http")
		if index == -1 {
			urlPath = fmt.Sprintf("%s/%s", baseUrl, path)
			oPath = fmt.Sprintf("%s%s", outPath, path)
		} else {
			oPath = fmt.Sprintf("%s/images/%s", outPath, filepath.Base(path))
		}
		ret, err = http.Get(urlPath, nil)
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}
		content := string(ret)
		createFile(oPath, content)
	}

	allUrl := getAllUrl(content)

	for _, u := range allUrl {
		subUrl := fmt.Sprintf("%s%s", baseUrl, u)
		if _, ok := pages[subUrl]; !ok {
			pages[subUrl] = subUrl
			go getPageAllData(outPath, baseUrl, subUrl)
		}
		fmt.Printf("%v\n", subUrl)
	}

}
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func createFile(path, content string) {
	paths, _ := filepath.Split(path)
	if !IsExist(paths) {
		os.MkdirAll(paths, 0755)
	}
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer file.Close()
	file.WriteString(content)
}
func getAllJavascript(content string) []string {
	re := regexp.MustCompile(`<script type="text/javascript" src="(.*?)"></script>`)
	finds := re.FindAllStringSubmatch(content, -1)
	var find []string
	for _, f := range finds {
		find = append(find, f[1])
	}
	return find
}
func getAllCss(content string) []string {
	re := regexp.MustCompile(`<link rel="stylesheet" type="text/css" href="(.*?)" />`)
	finds := re.FindAllStringSubmatch(content, -1)
	var find []string
	for _, f := range finds {
		find = append(find, f[1])
	}
	return find
}
func getAllUrl(content string) []string {
	re := regexp.MustCompile(`<a href="(/.*?[\d+])".*?>.*?</a>`)
	finds := re.FindAllStringSubmatch(content, -1)
	var find []string
	for _, f := range finds {
		find = append(find, f[1])
	}
	return find
}
func getAllImg(content string) []string {
	re := regexp.MustCompile(`<img src="(/.*?)">`)
	finds := re.FindAllStringSubmatch(content, -1)
	var find []string
	for _, f := range finds {
		find = append(find, f[1])
	}
	re = regexp.MustCompile(`background:#[\s\d]{6} url\((../images/[\s\S][.jpg|.png])\)`)
	finds = re.FindAllStringSubmatch(content, -1)
	for _, f := range finds {
		find = append(find, f[1])
	}
	return find
}
