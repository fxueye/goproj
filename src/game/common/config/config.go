package config

import (
	"bufio"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

// load configure file with "json" or "xml"
// "json" type will ignore '#' as comments at the end of line
func LoadConfig(fileType string, confPath string, config interface{}) (err error) {
	if err != nil {
		return
	}
	switch fileType {
	case "json":
		err = loadJsonFile(confPath, config)
	case "xml":
		err = loadXmlFile(confPath, config)
	default:
		err = errors.New("unsupported file type")
	}
	return
}

func loadJsonFile(filePath string, config interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bs := make([]byte, 0, 4096)
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			return err
		}
		// ignore comments by '#'
		i := bytes.IndexByte(line, '#')
		if i >= 0 {
			line = line[:i]
		}
		bs = append(bs, line...)
		if err == io.EOF {
			break
		}
	}
	return json.Unmarshal(bs, config)
}

func loadXmlFile(filePath string, config interface{}) error {
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return xml.Unmarshal(bs, config)
}
