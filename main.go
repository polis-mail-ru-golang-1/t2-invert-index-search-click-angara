package main

import (
	"bufio"
	"fmt"
	"github.com/click-angara/t2-invert-index-search-click-angara/invertedindex"
	"io/ioutil"
	"os"
	"strings"
)

// прочитали файл → понизили регистр → вычленили слова
func OperationOneFile(inxFile int, workFile []string) []string {

	test1, err := ioutil.ReadFile(workFile[inxFile]) // прочитали из файла
	if err != nil {
		panic(err)
	}
	test2 := strings.ToLower(string(test1)) // привели все к нижнему регитру

	var data []byte

	for i := 0; i < len(test2); i += 1 {
		if (test2[i] >= 'a' && test2[i] <= 'z') || test2[i] == '-' || test2[i] == ' ' || test2[i] == '\n' {
			if test2[i] == '\n' {
				data = append(data, ' ')
			} else {
				data = append(data, test2[i])
			}
		}
	} // выделили только слова

	datastr := strings.Split(string(data), " ")
	return datastr
}

func main() {
	var text []string
	flagMap := invertedindex.AddMap() // выделили память для структур
	if !flagMap {
		fmt.Println("ERR")
		os.Exit(1)
	}

	nameFiles := os.Args[1:]
	if len(nameFiles) < 1 {
		fmt.Println("Wrong format")
	} // Проверили корректность ввода

	// прочитали файл → понизили регистр → вычленили только слова
	for i, namef := range nameFiles {
		text = OperationOneFile(i, nameFiles)
		flagCheck := invertedindex.AddNewFile(text, namef)
		if !flagCheck {
			fmt.Println("Unable to add file:")
			os.Exit(1) // выходим из программы
		}
	}
	// Читаем запрос с консоли
	fmt.Print("Enter text: ")
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	text1 := myscanner.Text()

	line := strings.ToLower(text1)      // привели все к нижнему регитру
	example := strings.Split(line, " ") // преобразовали в массив строк
	example[len(example)-1] = strings.Trim(example[len(example)-1], "\n")

	var ForPrintSt []invertedindex.ForPrint
	ForPrintSt = (invertedindex.FileSearch(example)) // Ищем запрос в словаре
	for _, indx := range ForPrintSt {
		fmt.Println("-", indx.Namefi, "; совпадений -", indx.Count)
	}

}
