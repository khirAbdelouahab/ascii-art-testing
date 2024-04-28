package Data

import(
	"os"
	"log"
	"bufio"
)
// this function return the All Charachters From File

func LoadDataFromFile(fileName string)[95][8]string {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	
		scanner := bufio.NewScanner(file)
		var array [95][8]string
		for row := 1 ; row <= 95 ; row++ {
			col := 1
			for scanner.Scan() && col <= 8 {								
				line := scanner.Text() 
				if line != "" {                                                           //
					array[row-1][col-1] = line
					col++
				}
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		return array
}