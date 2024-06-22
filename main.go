package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args // Get command-line arguments
	if len(args) == 2 {
		ascii(args) // Call ascii function for simple usage
	} else if len(args) == 3 {
		ascii_fs(args) // Call ascii_fs function for file system output
	} else if len(args) == 4 {
		ascii_output(args) // Call ascii_output function for output to a specified file
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
	}
}

func ascii(args []string) {
	txt := args[1]

	textSlice := strings.Split(txt, "\\n") // Split text by "\n" for multiline support

	if !charValidation(txt) {
		fmt.Println("Error : invalid char")
		os.Exit(1)
	}

	file, err := os.ReadFile("standard.txt") // Read banner font file
	if err != nil {
		fmt.Println("Error : reading file")
		os.Exit(1)
	}
	slice := strings.Split(string(file), "\n") // Split font file into lines

	for j, txt := range textSlice {
		if txt != "" {
			for i := 0; i < 8; i++ { // Each letter is 8 lines high
				for _, v := range txt {
					firstLine := int(v-32)*9 + 1 + i // Determine the starting line for each character
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

	outputPtr := flag.String("output", "", "Output file name") // Define output file flag
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

	file, err := os.ReadFile(format + ".txt") // Read banner font file based on format
	if err != nil {
		fmt.Println("Error : reading file")
		os.Exit(1)
	}
	slice := strings.Split(string(file), "\n") // Split font file into lines

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
	os.WriteFile(*outputPtr, []byte(str), 0o644) // Write output to specified file
}

func ascii_fs(args []string) {
	txt := args[1]
	format := args[2]
	textSlice := strings.Split(txt, "\\n")

	if !charValidation(txt) {
		fmt.Println("Error : invalid char")
		os.Exit(1)
	}

	file, err := os.ReadFile(format + ".txt") // Read banner font file based on format
	if err != nil {
		fmt.Println("Error : reading file")
		os.Exit(1)
	}
	slice := strings.Split(string(file), "\n") // Split font file into lines

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

func charValidation(str string) bool {
	slice := []rune(str)
	for _, char := range slice {
		if char < 32 || char > 126 { // Validate characters to ensure they are within printable ASCII range
			return false
		}
	}
	return true
}
