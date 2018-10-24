package main

import (
	"./ForFunc"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	nameFiles := os.Args[1:]                         // Считываем названия файлов и разделяем их в массив строк
	vocabulary := make(map[string]map[string]int, 1) // слово → имя файла → количество вхождений
	fileList := make(map[string]map[string]int, 1)   // имя файла → слово → количество вхождений

	for indxfile := 0; indxfile < len(nameFiles); indxfile += 1 { // По очереди открываем файлы
		// Открываем файл и проверям корректность открытия
		infoFile, err := os.Open(string(nameFiles[indxfile]))
		if err != nil { // если возникла ошибка
			fmt.Println("Unable to create file:", err)
			os.Exit(1) // выходим из программы
		}
		defer infoFile.Close() // закрываем файл в конце работы

		// Начинаем читать файл по словам
		scanner := bufio.NewScanner(infoFile)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			var line string       // Здесь лежит считанная строка
			line = scanner.Text() // Считали строку из файла
			line = strings.Trim(line, "\n")

			if _, ok := vocabulary[line]; ok {
				vocabulary[line][nameFiles[indxfile]] += 1
			} else {
				tmp := make(map[string]int)
				tmp[nameFiles[indxfile]] = 1
				vocabulary[line] = tmp
			}
			if _, ok := fileList[nameFiles[indxfile]]; ok {
				fileList[nameFiles[indxfile]][line] += 1
			} else {
				tmp := make(map[string]int)
				tmp[line] = 1
				fileList[nameFiles[indxfile]] = tmp
			}
		}
	}

	// Читаем запрос с консоли
	fmt.Print("Enter text: ")
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	text := myscanner.Text()

	example := strings.Split(text, " ") // преобразовали в массив строк
	example[len(example)-1] = strings.Trim(example[len(example)-1], "\n")

	inst := 0
	for word, tmp := range vocabulary { // Ищем первое слово в словаре
		// !!! Что делать с пустым запросом
		if word == example[0] { // Нашли файл с совпадением
			for file, count := range tmp { // проходим поочередно по всем файлам, в которых есть первое слово
				// ищем каждый файл во вспомогательной map
				countwrite := count
				flag := 1
				for file1, _ := range fileList {
					if file1 == file { // значит ищем в этом файле все введенные слова
						for indx := 1; indx < len(example); indx += 1 {
							for word1, count1 := range fileList[file1] {
								if example[indx] == word1 {
									countwrite += count1
									flag += 1
									break
								}
							}
						}
						if flag == len(example) { // значит встретили все слова
							flag = ForFunc.AddStruct(countwrite, file)
							inst += 1
						}
					}
				}
			}
			break
		}
	}
	ForFunc.SortStruct()
}
