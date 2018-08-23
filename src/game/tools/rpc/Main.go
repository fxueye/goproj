package main

import (
	"flag"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type RPCData struct {
	ClzName string
	HasWrap bool
	Datas   []*RPCDataElem
}
type RPCDataElem struct {
	Opcode  int
	OpDef   string
	Func    string
	Args    []*RPCDataArg
	Comment string
}
type RPCDataArg struct {
	RpcType  string
	RpcValue string
}

type WrapData struct {
	Name     string
	Values   []string
	Types    []string
	Repeats  []bool
	Comments []string
}

var (
	inputPath   string
	outputPath  string
	tmpRpcPath  string
	tmpWrapPath string
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

	tmpRpcPath, _ = filepath.Abs(*temp_path)
	fmt.Printf("template_rpc_path=%s\n", tmpRpcPath)

	tmpWrapPath = fmt.Sprintf("%s%c%s", tmpRpcPath, os.PathSeparator, "wrap")
	fmt.Printf("template_wrap_path=%s\n", tmpWrapPath)

	if !IsExist(inputPath) {
		panic(fmt.Sprint("input path not exists: %s", inputPath))
	}
	if !IsExist(tmpRpcPath) {
		panic(fmt.Sprint("tmpRpcPath path not exists: %s", tmpRpcPath))
	}
	if !IsExist(tmpWrapPath) {
		panic(fmt.Sprint("tmpWrapPath path not exists: %s", tmpWrapPath))
	}

	os.MkdirAll(outputPath, 0755)
	outputWrapPath := fmt.Sprintf("%s%c%s", outputPath, os.PathSeparator, "wrap")
	os.MkdirAll(outputWrapPath, 0755)

	CreateRPCFiles()
	CreateWrapFiles()
}

func CreateRPCFiles() {
	dir, _ := ioutil.ReadDir(inputPath)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		ext := filepath.Ext(fi.Name())
		name := fi.Name()
		clzName := name[:strings.Index(name, ".")]
		if name == "" || name[0] == '~' || name[0] == '.' {
			continue
		}
		if ext != ".xls" && ext != ".xlsx" {
			continue
		}

		data := ParseRPCFile(fmt.Sprintf("%s%c%s", inputPath, os.PathSeparator, fi.Name()), clzName)
		if(strings.Index(clzName,"Client") >= 0){
			CreateFile(fmt.Sprintf("%s%c%sCodes.cs", outputPath, os.PathSeparator, clzName),
			fmt.Sprintf("%s%c%s", tmpRpcPath, os.PathSeparator, "tmp_cs_opcode.txt"), data)
		}
		
		CreateFile(fmt.Sprintf("%s%c%sCodes.go", outputPath, os.PathSeparator, clzName),
			fmt.Sprintf("%s%c%s", tmpRpcPath, os.PathSeparator, "tmp_go_opcode.txt"), data)

		if strings.Index(clzName, "Client") >= 0 {
			CreateFile(fmt.Sprintf("%s%c%sInvoker.cs", outputPath, os.PathSeparator, clzName),
				fmt.Sprintf("%s%c%s", tmpRpcPath, os.PathSeparator, "tmp_cs_invoker.txt"), data)
			CreateFile(fmt.Sprintf("%s%c%sInvoker.go", outputPath, os.PathSeparator, clzName),
				fmt.Sprintf("%s%c%s", tmpRpcPath, os.PathSeparator, "tmp_go_invoker.txt"), data)
		}

		if strings.Index(clzName, "Server") >= 0 {
			CreateFile(fmt.Sprintf("%s%c%sInvoker.go", outputPath, os.PathSeparator, clzName),
				fmt.Sprintf("%s%c%s", tmpRpcPath, os.PathSeparator, "tmp_go_invoker.txt"), data)
		}
	}
}

func CreateWrapFiles() {
	input := fmt.Sprintf("%s%c%s", inputPath, os.PathSeparator, "Wrap")
	dir, _ := ioutil.ReadDir(input)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		name := fi.Name()
		ext := filepath.Ext(name)
		clzName := name[:strings.Index(name, ".")]
		if name == "" || name[0] == '~' || name[0] == '.' {
			continue
		}
		if ext != ".xls" && ext != ".xlsx" {
			continue
		}

		data := ParseWrapFile(fmt.Sprintf("%s%c%s", input, os.PathSeparator, fi.Name()), clzName)
		CreateFile(fmt.Sprintf("%s%c%s%c%s.cs", outputPath, os.PathSeparator, "wrap", os.PathSeparator, clzName),
			fmt.Sprintf("%s%c%s", tmpWrapPath, os.PathSeparator, "tmp_cs.txt"), data)

		CreateFile(fmt.Sprintf("%s%c%s%c%s.go", outputPath, os.PathSeparator, "wrap", os.PathSeparator, clzName),
			fmt.Sprintf("%s%c%s", tmpWrapPath, os.PathSeparator, "tmp_go.txt"), data)
	}

}

func CreateFile(path string, tmpPath string, data interface{}) {
	fmt.Println(path)
	fmt.Println(tmpPath)
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

func ParseWrapFile(path string, clzName string) *WrapData {
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		panic(fmt.Sprint("cannot open file: ", path))
	}
	sheet := xlFile.Sheets[0]

	data := &WrapData{
		clzName,
		make([]string, 0, 16),
		make([]string, 0, 16),
		make([]bool, 0, 16),
		make([]string, 0, 16),
	}
	for r, row := range sheet.Rows {
		if r < 1 {
			continue
		}
		n := row.Cells[0].Value
		t := row.Cells[1].Value
		r := row.Cells[2].Bool()
		c := ""
		if len(row.Cells) > 3 {
			c = row.Cells[3].Value
		}
		data.Values = append(data.Values, n)
		data.Types = append(data.Types, t)
		data.Repeats = append(data.Repeats, r)
		data.Comments = append(data.Comments, c)
	}
	return data
}

func ParseRPCFile(path string, clzName string) *RPCData {
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		panic(fmt.Sprint("cannot open file: ", path))
	}
	sheet := xlFile.Sheets[0]

	data := &RPCData{
		clzName,
		false,
		make([]*RPCDataElem, 0, 32),
	}
	for r, row := range sheet.Rows {
		if r < 1 {
			continue
		}
		opcode, err := row.Cells[0].Int()
		if err != nil {
			fmt.Printf("%v\n",clzName)
			panic(err)
		}
		opdef := row.Cells[1].Value
		fnName := row.Cells[2].Value

		args := make([]*RPCDataArg, 0, 4)
		argStr := strings.TrimSpace(row.Cells[3].Value)
		if argStr != "" {
			strs := strings.Split(argStr, ",")
			for _, str := range strs {
				ss := strings.Split(strings.TrimSpace(str), " ")
				if len(ss) != 2 {
					panic(fmt.Sprint("invalid args, row=", ss))
				}
				if isWrap(ss[0]) {
					data.HasWrap = true
				}
				args = append(args, &RPCDataArg{ss[0], ss[1]})
			}
		}
		comment := row.Cells[4].Value
		data.Datas = append(data.Datas, &RPCDataElem{opcode, opdef, fnName, args, comment})
	}

	return data
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func getGoType(t string) string {
	switch t {
	case "short":
		return "int16"
	case "int":
		return "int32"
	case "long":
		return "int64"
	case "float":
		return "float32"
	case "double":
		return "float64"

	default:
		return t
	}

}
func (*RPCData) GetGoType(t string) string {
	return getGoType(t)
}

func (*WrapData) GetGoType(t string) string {
	return getGoType(t)
}

func (*RPCDataArg) GetCSFunc(t string, isGet bool) string {
	return getCSFunc(t, isGet)
}
func (*WrapData) GetCSFunc(t string, isGet bool) string {
	return getCSFunc(t, isGet)
}
func getCSFunc(t string, isGet bool) string {
	switch t {
	case "bool":
		if isGet {
			return "GetBool"
		} else {
			return "PutBool"
		}
	case "short":
		if isGet {
			return "GetShort"
		} else {
			return "PutShort"
		}
	case "int":
		if isGet {
			return "GetInt"
		} else {
			return "PutInt"
		}
	case "long":
		if isGet {
			return "GetLong"
		} else {
			return "PutLong"
		}
	case "float":
		if isGet {
			return "GetFloat"
		} else {
			return "PutFloat"
		}
	case "double":
		if isGet {
			return "GetDouble"
		} else {
			return "PutDouble"
		}
	case "string":
		if isGet {
			return "GetString"
		} else {
			return "PutString"
		}
	default:
		return ""
	}
}

func (*WrapData) GetGoFunc(t string, isGet bool) string {
	return getGoFunc(t, isGet)
}
func (*RPCDataArg) GetGoFunc(t string, isGet bool) string {
	return getGoFunc(t, isGet)
}

func getGoFunc(t string, isGet bool) string {
	switch t {
	case "bool":
		if isGet {
			return "PopBool"
		} else {
			return "PutBool"
		}
	case "short":
		if isGet {
			return "PopInt16"
		} else {
			return "PutInt16"
		}
	case "int":
		if isGet {
			return "PopInt32"
		} else {
			return "PutInt32"
		}
	case "long":
		if isGet {
			return "PopInt64"
		} else {
			return "PutInt64"
		}
	case "float":
		if isGet {
			return "PopFloat32"
		} else {
			return "PutFloat32"
		}
	case "double":
		if isGet {
			return "PopFloat64"
		} else {
			return "PutFloat64"
		}
	case "string":
		if isGet {
			return "PopString"
		} else {
			return "PutString"
		}
	default:
		return ""
	}
}

func (*RPCDataArg) IsWrap(t string) bool {
	return isWrap(t)
}

func (*WrapData) IsWrap(t string) bool {
	return isWrap(t)
}

func isWrap(t string) bool {
	switch t {
	case "bool", "short", "int", "long", "float", "double", "string":
		return false
	default:
		return true

	}
}
