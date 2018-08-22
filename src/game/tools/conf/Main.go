package main

import (
	"flag"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

var (
	inputPath       string
	outputPath      string
	tmpJsonPath     string
	tmpCSPath       string
	tmpCSLoaderPath string
	tmpGoPath       string
	tmpGoLoaderPath string
	translatePath   string
)

func main() {
	input_path := flag.String("I", "", "configure input path")
	output_path := flag.String("O", "", "configure output path")
	temp_path := flag.String("T", "", "configure template path")
	flag.Parse()

	inputPath, _ = filepath.Abs(*input_path)
	fmt.Printf("input_path=%s\n", inputPath)

	outputPath, _ = filepath.Abs(*output_path)
	fmt.Printf("output_path=%s\n", outputPath)

	tmpJsonPath, _ = filepath.Abs(fmt.Sprintf("%s%c%s", *temp_path, os.PathSeparator, "tmp_json.txt"))
	fmt.Printf("temp_json_path=%s\n", tmpJsonPath)

	tmpCSPath, _ = filepath.Abs(fmt.Sprintf("%s%c%s", *temp_path, os.PathSeparator, "tmp_cs.txt"))
	fmt.Printf("temp_csharp_path=%s\n", tmpCSPath)

	tmpCSLoaderPath, _ = filepath.Abs(fmt.Sprintf("%s%c%s", *temp_path, os.PathSeparator, "tmp_cs_loader.txt"))
	fmt.Printf("temp_csharp_loader_path=%s\n", tmpCSLoaderPath)

	tmpGoPath, _ = filepath.Abs(fmt.Sprintf("%s%c%s", *temp_path, os.PathSeparator, "tmp_go.txt"))
	fmt.Printf("temp_json_path=%s\n", tmpJsonPath)

	tmpGoLoaderPath, _ = filepath.Abs(fmt.Sprintf("%s%c%s", *temp_path, os.PathSeparator, "tmp_go_loader.txt"))
	fmt.Printf("temp_go_loader_path=%s\n", tmpGoLoaderPath)

	translatePath, _ = filepath.Abs(fmt.Sprintf("%s%c%s", outputPath, os.PathSeparator, "LangConf.txt"))
	fmt.Printf("translate_text_path=%s\n", translatePath)

	if !IsExist(inputPath) {
		panic(fmt.Sprint("input path not exists: %s", inputPath))
	}
	if !IsExist(tmpJsonPath) {
		panic(fmt.Sprint("json template path not exists: %s", tmpJsonPath))
	}
	if !IsExist(tmpCSPath) {
		panic(fmt.Sprint("csharp template path not exists: %s", tmpCSPath))
	}
	if !IsExist(tmpCSLoaderPath) {
		panic(fmt.Sprint("csharp loader template path not exists: %s", tmpCSLoaderPath))
	}
	if !IsExist(tmpGoPath) {
		panic(fmt.Sprint("go template path not exists: %s", tmpGoPath))
	}
	if !IsExist(tmpGoLoaderPath) {
		panic(fmt.Sprint("go loader template path not exists: %s", tmpGoLoaderPath))
	}
	if outputPath == "" {
		panic(fmt.Sprint("invalid output path: %s", outputPath))
	}
	if IsExist(outputPath) {
		os.RemoveAll(outputPath)
	}
	jsonPath := fmt.Sprintf("%s%c%s", outputPath, os.PathSeparator, "json")
	csPath := fmt.Sprintf("%s%c%s", outputPath, os.PathSeparator, "cs")
	goPath := fmt.Sprintf("%s%c%s", outputPath, os.PathSeparator, "go")
	os.MkdirAll(jsonPath, 0755)
	os.MkdirAll(csPath, 0755)
	os.MkdirAll(goPath, 0755)

	confNames := make([]string, 0, 32)
	textMap := make(map[string]bool)
	fmt.Printf("build config :-------------------------\n")
	dir, _ := ioutil.ReadDir(inputPath)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}

		name := fi.Name()
		fmt.Printf("build config: %s\n", name)

		ext := filepath.Ext(name)
		clzName := name[:strings.Index(name, ".")]

		if name == "" || name[0] == '~' || name[0] == '.' {
			continue
		}
		if ext != ".xls" && ext != ".xlsx" {
			continue
		}

		conf := ParseFile(fmt.Sprintf("%s%c%s", inputPath, os.PathSeparator, fi.Name()), clzName, textMap)
		CreateFile(fmt.Sprintf("%s%c%s.txt", jsonPath, os.PathSeparator, clzName), tmpJsonPath, conf)
		CreateFile(fmt.Sprintf("%s%c%s.cs", csPath, os.PathSeparator, clzName), tmpCSPath, conf)
		CreateFile(fmt.Sprintf("%s%c%s.go", goPath, os.PathSeparator, clzName), tmpGoPath, conf)

		confNames = append(confNames, clzName)
	}

	CreateFile(fmt.Sprintf("%s%c%s.go", goPath, os.PathSeparator, "ConfLoader"), tmpGoLoaderPath, confNames)
	CreateFile(fmt.Sprintf("%s%c%s.cs", csPath, os.PathSeparator, "ConfLoader"), tmpCSLoaderPath, confNames)
	fmt.Println("----------------------------------\n")

	CreateTranslateFile(textMap, translatePath)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

type ConfData struct {
	ClzName       string
	ConfNames     []string
	ConfTypes     []string
	ConfIsArrays  []bool
	ConfDatas     [][]string
	NeedTranslate []bool
}

func CreateTranslateFile(textMap map[string]bool, path string) {
	f, _ := os.Create(path)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	texts := make([]string, 0, len(textMap))
	for k, _ := range textMap {
		texts = append(texts, k)
	}
	sort.Strings(texts)

	for _, k := range texts {
		f.WriteString(fmt.Sprintln(k))
	}
}

func CreateFile(path string, tmpPath string, data interface{}) {
	f, _ := os.Create(path)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()

	t, err := template.ParseFiles(tmpPath)
	if err != nil {
		panic(err)
	}
	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}
}

func ParseFile(path string, clzName string, textMap map[string]bool) *ConfData {
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		panic(fmt.Sprint("cannot open file: ", path))
	}
	sheet := xlFile.Sheets[0]

	data := &ConfData{
		clzName,
		make([]string, 0, 32),
		make([]string, 0, 32),
		make([]bool, 0, 32),
		make([][]string, 0, 32),
		make([]bool, 0, 32),
	}
	for _, cell := range sheet.Rows[0].Cells {
		v := cell.Value
		if v == "" {
			break
		}

		data.NeedTranslate = append(data.NeedTranslate, v[0] == 'T')
	}
	for _, cell := range sheet.Rows[1].Cells {
		v := cell.Value
		if v == "" {
			break
		}

		data.ConfNames = append(data.ConfNames, fmt.Sprint(strings.ToUpper(v[:1]), v[1:]))
	}
	for _, cell := range sheet.Rows[2].Cells {
		v := cell.Value
		if v == "" {
			break
		}
		if strings.Index(cell.Value, "[]") > 0 {
			data.ConfIsArrays = append(data.ConfIsArrays, true)
			data.ConfTypes = append(data.ConfTypes, cell.Value[:strings.Index(cell.Value, "[]")])
		} else {
			data.ConfIsArrays = append(data.ConfIsArrays, false)
			data.ConfTypes = append(data.ConfTypes, cell.Value)
		}

	}
	if len(data.ConfNames) != len(data.ConfTypes) {
		panic("configure len(names) != len(types)")
	}

	length := len(data.ConfNames)
	for r, row := range sheet.Rows {
		if r < 3 {
			continue
		}
		if len(row.Cells) == 0 || row.Cells[0] == nil || row.Cells[0].Value == "" {
			break
		}
		d := make([]string, length)
		for i := 0; i < length; i++ {
			if data.ConfIsArrays[i] {
				if  i < len(row.Cells) && row.Cells[i].Value != "" {
					strs := strings.Split(row.Cells[i].Value, ",")
					ss := make([]string, len(strs))
					for j := 0; j < len(strs); j++ {
						ss[j] = GetValueText(strs[j], data.ConfTypes[i])
						if data.NeedTranslate[i] {
							textMap[strs[j]] = true
						}
					}
					d[i] = strings.Join(ss, ",")
				} else {
					d[i] = ""
				}
			} else {
				cellValue := ""
				if (i < len(row.Cells)) {
					cellValue = row.Cells[i].Value
				}
				d[i] = GetValueText(cellValue, data.ConfTypes[i])
				if data.NeedTranslate[i] {
					textMap[row.Cells[i].Value] = true
				}
			}
		}
		data.ConfDatas = append(data.ConfDatas, d)
	}

	return data
}

func GetValueText(v string, t string) string {
	switch t {
	case "bool":
		if v != "0" && v != "FALSE" && v != "false" {
			return "true"
		}
		return "false"
	case "string":
		v = strings.Replace(v, "\\", "\\\\", -1)
		v = strings.Replace(v, "\"", "\\\"", -1)
		return fmt.Sprint("\"", v, "\"")
	case "float":
		if v == "" {
			return ""
		}
		if strings.Index(v, ".") < 0 {
			return fmt.Sprint(v, ".0")
		}
		return v
	case "int":
		if v == "" {
			return "0"
		}
		return v
	default:
		return v
	}
}
