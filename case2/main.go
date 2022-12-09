package main

import (
	"fmt"
	"path/filepath"

	"github.com/frtatmaca/case2/helper"
)

func main() {
	name := "data/wordlist.txt"
	absFname, err := filepath.Abs(name)

	lines, err := helper.ReadLines(absFname)
	if err != nil {
		fmt.Println("Error: %s\n", err)
		return
	}

	anagrams := ""
	for _, a := range lines {
		anagrams = ""
		for _, b := range lines {
			if a == b {
				continue
			}
			if helper.IsAnagram(a, b) {
				anagrams = anagrams + b + ","
				lines = helper.Remove(lines, b)
			}
		}

		if len(anagrams) > 0 {
			anagrams = a + "," + anagrams
			fmt.Println(anagrams)
		}

		lines = helper.Remove(lines, a)
	}
}
