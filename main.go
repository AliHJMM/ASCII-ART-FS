package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		ascii(args)
	} else if len(args) == 3 {
		ascii_fs(args)
	} else if len(args) == 4 {
		ascii_output(args)
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standar")
	}
}

func ascii(args []string) {
	txt := args[1]

	textSlice := strings.Split(txt, "\\n")

	if !charValidation(txt) {
		fmt.Println("Error : invalid char")
		os.Exit(1)
	}
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error : reading file")
		os.Exit(1)
	}
	slice := strings.Split(string(file), "\n")
	for j, txt := range textSlice {
		if txt != "" {
			for i := 0; i < 8; i++ {
				for _, v := range txt {
					firstLine := int(v-32)*9 + 1 + i
					fmt.Print(slice[firstLine])
				}
				fmt.Println()
			}
		} else if j != len(textSlice)-1 {
			fmt.Println("")
		}
	}
}

func ascii_output(args []string) {
	txt := args[2]

	format := args[3]
	str := ""
	outputPtr := flag.String("output", "", "Output file name")
	flag.Parse()

	if *outputPtr == "" {
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		os.Exit(1)
	}

	textSlice := strings.Split(txt, "\\n")

	if !charValidation(txt) {
		fmt.Println("Error : invalid char")
		os.Exit(1)
	}
	file, err := os.ReadFile(format + ".txt")
	if err != nil {
		fmt.Println("Error : reading file")
		os.Exit(1)
	}
	slice := strings.Split(string(file), "\n")
	for j, txt := range textSlice {
		if txt != "" {
			for i := 0; i < 8; i++ {
				for _, v := range txt {
					firstLine := int(v-32)*9 + 1 + i
					str += slice[firstLine]
				}
				str += "\n"
			}
		} else if j != len(textSlice)-1 {
			str += "\n"
		}
	}
	os.WriteFile(*outputPtr, []byte(str), 0o644)
}

