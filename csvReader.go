package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//       "math"
	//      "runtime"
)

func rec(p int) int {
	if p < 0 {
		return 0
	} else {
		return p + (rec(p - 1))
	}
}

func main() {
	fmt.Println("CSV Reader, JTM, 2018")
	// fmt.Println("jetzt auch mit git und vscode")
	fmt.Println("Tolle Sache, dieser VSCode")
	// toll

	// show runtime arguments
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	// check for filenames

	files := os.Args[1:]

	if len(files) == 0 {
		// keine Dateien angegeben
	} else {
		for _, arg := range files {

			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			// process csv/json data
			processLines(f)

			f.Close()
		}
	}
}

func processLines(f *os.File) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		fmt.Println(input.Text())

		for _, field := range strings.Split(string(input.Text()), ";") {
			fmt.Println(field)
		}
	}
}
