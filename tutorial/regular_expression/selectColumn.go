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
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("usage: selectColumn column <file1> [<file2> [...<fileN>]]\n")
		os.Exit(1)
	}
	temp, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Column value is not an inteager:", temp)
		return
	}

	column := temp
	if column < 0 {
		fmt.Println("Invalid Column number")
		os.Exit(1)
	}

	for _, filename := range args[2:] {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			continue
		}
		defer f.Close()
		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n') // 매개변수로 지정한 내용을 처음 발견할 때까지 파일을 읽음
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
			}
			data := strings.Fields(line) // 빈칸을 기준으로 문장을 나눔
			if len(data) >= column {
				fmt.Println((data[column-1]))
			}
		}
	}
}
