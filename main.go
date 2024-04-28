package main

import(
	"ascii-art/Data"
    "fmt"
    "os"
	"strings"
)
func PrintCharachters(s, FromFile string) { // "Hello\n\nworld" // "\n\n\n"   // "" // "hello\n"
	size_s := len(s)
	activeIndex := 0
	line := 1
	startIndex := 0
	Charachters := Data.LoadDataFromFile(FromFile)
	for activeIndex < size_s || line <= 8 {
		if s[activeIndex] == '\n' && line < 8{
			activeIndex = startIndex ; line++
		}
		indexChar := FindCharacterId(rune(s[activeIndex]))
		PrintCharLine(Charachters,indexChar,line-1)
		activeIndex++; line++
	}
}
func ScanCharachters(s string) {
	size_s := len(s)
	if size_s == 0 {
		return 
	}
	Charachters := Data.LoadDataFromFile("standard.txt")
	for line := 1; line <= 8 ; line++ {
		for i := 0 ; i <= size_s - 1; i++ {
			indexChar := FindCharacterId(rune(s[i]))
			if indexChar == -1 {
				fmt.Println()
				continue
			}
			PrintCharLine(Charachters,indexChar,line-1)
		}
		fmt.Println()
	}
}

func FindCharacterId(c rune) int {
	index := int(c - 32)
	if index < 0 {
		return -1
	}
	if index >= 95 {
		return -2
	} 
	return index
}
func PrintCharLine(ArrayOfCharchters [95][8]string, index,line int) {
	fmt.Printf("%v",ArrayOfCharchters[index][line])
}

func Ascii_Art(input string){
	if input == "" {
		return 
	}
	myslice := strings.Split(input,"\\n")
	for _, str := range myslice {
		if str == "" {
			fmt.Println()
			continue
		}
		ScanCharachters(str)
	}
}
func main()  {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("ERROR: Enter a valid argument!")
		return
	}
	Ascii_Art(arguments[1])
}