package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//Email  is
var Email map[string][]string

func main() {
	Email = make(map[string][]string)

	for _, file := range os.Args[1:] {
		fi, err := os.Open(file)
		if err != nil {
			fmt.Println("open file error ", err)
		}

		br := bufio.NewReader(fi)
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			if contains(string(a)) {
				continue
			}

			sp := strings.Split(string(a), "|")
			email := sp[4]
			if email != "null" {
				Email[email] = sp
			}
		}
	}

	for k, v := range Email {
		//TODO send email here
		fmt.Println(k, "|", v)
	}
}

func contains(a string) bool {
	return strings.Contains(a, "iflytek") || strings.Contains(a, "Hefei") || strings.Contains(a, "hefei")
}
