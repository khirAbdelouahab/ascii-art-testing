package main
import ("ascii-art/Data" ; "fmt" ; "os" ; "strings")
func PrintCharachters(s, FromFile string) { // "\nHello\n\nworld" // "\n\n\n"   // "" // "hello\n"
	size_s := len(s) ; activeIndex := 0 ; line := 1 ; startIndex := 0 ; Charachters := Data.LoadDataFromFile(FromFile)
	for activeIndex < size_s {
		if s[activeIndex] != '\n' {
			if indexChar := FindCharacterId(rune(s[activeIndex])) ; indexChar != -1 {
				PrintCharLine(Charachters,indexChar,line-1)
				activeIndex++;
				if activeIndex == size_s || (activeIndex == size_s - 1 && s[activeIndex]=='\n') {
					if line < 8 {
						line++ ; activeIndex = startIndex
					}
					fmt.Println()
				}
			} else {
				activeIndex++
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
	if index < 0 || index >= 95 {return -1}
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