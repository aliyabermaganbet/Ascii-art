package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		t := "standard.txt"
		if Check_for_hash(t) {
			file, err := os.Open(t)
			if err != nil {
				log.Fatal(err)
			}
			r := bufio.NewScanner(file)
			var array []string
			for r.Scan() {
				array = append(array, r.Text())
			}
			myMap := make(map[rune][]string)
			var q rune = 32
			for i := 1; i < len(array); i += 9 {
				myMap[q] = array[i : i+8]
				q++
			}
			arg := os.Args[1]
			if Check_for_letters(arg) {
				str := ""
				k := strings.ReplaceAll(arg, "\\n", "\n")
				rk := strings.Split(k, "\n")
				for _, rn := range rk {
					if rn == "" {
						str += "\n"
					} else {
						for i := 0; i < 8; i++ {
							for d := 0; d < len(rn); d++ {
								str += myMap[rune(rn[d])][i]
							}
							str += "\n"
						}
					}
				}
				fmt.Print(str)
			} else {
				fmt.Println("Invalid syntax")
			}
		}
	} else {
		return
	}
}

func Check_for_letters(s string) bool {
	for i := 0; i < len(s); i++ {
		if !(s[i] >= 32 && s[i] <= 126) {
			return false
		}
	}
	return true
}

func hash(s string) string {
	h := md5.New()
	f, err := os.Open(s)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	a := fmt.Sprintf("%x", h.Sum(nil))
	return a
}

func Check_for_hash(s string) bool {
	switch s {
	case "standard.txt":
		{
			if hash(s) == "ac85e83127e49ec42487f272d9b9db8b" {
				return true
			}
		}
	}
	return false
}
