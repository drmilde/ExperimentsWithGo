package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	//      "runtime"
)

type event struct {
	TimeStamp float64 `json:"timestamp,omitempty"`
	C1        float64 `json:"c1,omitempty"`
	C2        float64 `json:"c2,omitempty"`
}

func main() {
	fmt.Println("CSV Reader, JTM, 2018")

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
		//fmt.Println(input.Text())

		fields := strings.Split(string(input.Text()), ";")
		values := make([]float64, len(fields))

		i := 0
		for _, field := range fields {
			//fmt.Println(field)
			sfield := strings.Replace(field, ",", ".", 1)

			val, err := strconv.ParseFloat(sfield, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			} else {
				values[i] = val
				i++
			}
		}

		//ev := event(timeStamp: values[0], c1: values[1], c2: values[2])
		ev := event{values[0], values[1], values[2]}

		//nm := name{"Rebecca", "Milde"}
		js, _ := json.Marshal(ev)
		fmt.Printf("%s\n", js)
	}
}
