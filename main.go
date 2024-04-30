package main
import ("ascii-art/Data" ; "fmt" ; "os" ; "strings")
func PrintCharachters(s, FromFile string) { // "\nHello\n\nworld" // "\n\n\n"   // "" // "hello\n"
	size_s := len(s) ; activeIndex := 0 ; line := 1 ; startIndex := 0 ; Charachters := Data.LoadDataFromFile(FromFile)
	for activeIndex < size_s {
		if s[activeIndex] != '\n' {
			indexChar := FindCharacterId(rune(s[activeIndex]))
			PrintCharLine(Charachters,indexChar,line-1)
			activeIndex++;
			if activeIndex == size_s  && line < 8 {
				activeIndex = startIndex ; line++ ; fmt.Println()
			}
		} else {
			if startIndex == activeIndex || line == 8 {
				fmt.Println() ; startIndex = activeIndex + 1 ; activeIndex = startIndex ; line = 1
			} else if line < 8 {
				line++ ; activeIndex = startIndex ;fmt.Println()
			}
		}
	}
}
func FindCharacterId(c rune) int {
	index := int(c - 32)
	if index < 0 {return -1}
	if index >= 95 {return -2} 
	return index
}
func PrintCharLine(ArrayOfCharchters [95][8]string, index,line int) {
	fmt.Printf("%v",ArrayOfCharchters[index][line])
}
func main()  {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("ERROR: Enter a valid argument!")
		return
	}
	arguments[1] = strings.ReplaceAll(arguments[1],"\\n","\n") 
	PrintCharachters(arguments[1],"standard.txt")
}