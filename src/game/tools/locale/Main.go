package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"path/filepath"
)

var (
	inputPath  string
	outputPath string
)

func main() {
	input_path := flag.String("I", "", "locale input path")
	output_path := flag.String("O", "", "locale output path")
	flag.Parse()

	inputPath, _ = filepath.Abs(*input_path)
	fmt.Printf("input_path=%s\n", inputPath)

	outputPath, _ = filepath.Abs(*output_path)
	fmt.Printf("output_path=%s\n", outputPath)

	f, _ := os.Create(outputPath)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()

	locales := make([]string, 0, 4)
	if IsExist(inputPath) {
		xlFile, err := xlsx.OpenFile(inputPath)
		if err != nil {
			panic(fmt.Sprint("cannot open file: ", inputPath))
		}
		sheet := xlFile.Sheets[0]

		for _, cell := range sheet.Rows[0].Cells {
			v := cell.Value
			if v == "" {
				break
			}
			locales = append(locales, v)
		}

		for r, row := range sheet.Rows {
			if r < 1 {
				continue
			}
			if len(row.Cells) == 0 || row.Cells[0] == nil || row.Cells[0].Value == "" {
				break
			}
			data := make(map[string]string)
			for c, cell := range row.Cells {
				if c < len(locales) {
					data[locales[c]] = cell.Value
				}
			}

			bs, _ := json.Marshal(data)
			f.WriteString(fmt.Sprintln(string(bs)))
		}
	}
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
