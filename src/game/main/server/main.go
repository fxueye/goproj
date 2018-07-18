package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./update_mysql.txt")
	if err != nil {
		return
	}
	defer file.Close()
	wfile, err := os.Create("./update_mysql_.txt")
	if err != nil {
		return
	}
	defer wfile.Close()

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			return
		}
		strs := strings.Split(string(line), "	")
		if len(strs) == 2 {
			enum, _ := strconv.ParseFloat(strs[1], 64)
			estr := fmt.Sprintf("%e", enum)
			b := strs[0] == estr
			wstr := fmt.Sprintf("%s	%s	%s	%v\n", strs[0], strs[1], estr, b)
			if !b {
				fmt.Sprintf("not == :%v", wstr)
			}
			wfile.WriteString(wstr)
		} else {
			fmt.Sprintf("error line:%v", string(line))
		}

		// fmt.Println(string(line))
		if err == io.EOF {
			break
		}
	}

}
