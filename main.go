package main

import (
	"bufio"
	"fmt"
	"os"
)


// Структура для хранения слов и информации о них
type TableVocab struct {
	word string				// Место для слова
	count []int				// Сколько раз встретилось слово в файле
	namefile []string		// Название файла в котором встретилось слово
	indxfile int
}

// Сортируем по частоте встречаемости в файле
func Sort (array TableVocab, countelen int) TableVocab{

	for i:= 0; i < countelen - 1; i += 1{
		for j := i; j < countelen; j += 1{
			if array.count[i] < array.count[j]{
				tmp := array.count[i]
				array.count[i] = array.count[j]
				array.count[j] = tmp

				ptr :=array.namefile[i]
				array.namefile[i] = array.namefile[j]
				array.namefile[j] = ptr
			}
		}
	}

	return array
}
// При добавлении нового слова, заполняем всю структуру для него
func writeVocab (line string, name string, i int,len int ) TableVocab{
	var array TableVocab
	array.count = make ([]int,len)
	array.namefile = make ([]string,len)


	array.word = line
	array.count[i] = 1 					// Так как добавляем новое слово → всегда количество 1
	array.namefile[i] = name
	array.indxfile = i 					// Записали индекс файла
	return array
}

func main() {

	nameFiles := os.Args[1:] 				// Считываем названия файлов и разделяем их в массив строк
	vocabulary := make( []TableVocab,5) 	// В массиве храним структуры со словами и их встречаемостью
	countword := 0 							// количество записанных слов в словаре

	for i := 0; i < len(nameFiles); i += 1 {	// По очереди открываем файлы

		// Открываем файл и проверям корректность открытия
		infoFile,err := os.Open(string(nameFiles[i]))
		if err != nil{                          // если возникла ошибка
			fmt.Println("Unable to create file:", err)
			os.Exit(1)                          // выходим из программы
		}
		defer infoFile.Close()                      // закрываем файл в конце работы

		// Начинаем читать файл по словам
		scanner := bufio.NewScanner(infoFile)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			var line string					// Здесь лежит считанная строка
			line = scanner.Text()			// Считали строку из файла

			if countword == 0 { // Если нет слов в словаре
				vocabulary[0] = writeVocab (line,nameFiles[i],i, len(nameFiles) ) // Добавляем в массив структур элемент с новым словом
				countword += 1

			} else{ // Если в словаре уже есть слова
				readflag := false
				for j := 0; j < countword; j +=1 { 					// Ищем совпадения
					if vocabulary[j].word == line { 				// если нашли такое слово
						vocabulary[j].count[i] += 1					// Плюсуем количество
						vocabulary[j].namefile[i] = nameFiles[i]	// Записываем имя файла, где встетили
						readflag = true
					}
				}
				if(!readflag) {			// Если не нашли совпадений, добавляем как новое слово
					if countword < len(vocabulary) {
						vocabulary[countword] = writeVocab (line,nameFiles[i],i,len(nameFiles))
					} else {
						vocabulary = append(vocabulary, writeVocab (line,nameFiles[i],i,len(nameFiles)))
					}
					countword += 1
				}
			}
			//fmt.Println(vocabulary)
		}
	}

//  Считали строку для поиска
	fmt.Print("Введите строку для поиска : ")

	var text string
	fmt.Scanln(&text)

// Для одной строки
	for i := 0; i < countword; i += 1{
		if (vocabulary[i].word == text){
			forPrintStruct := Sort (vocabulary[i], len(nameFiles))
			t := 0
			for{
				if  t < len(forPrintStruct.namefile) && forPrintStruct.count[t] > 0  {
					fmt.Print("-", forPrintStruct.namefile[t],";	совпадений - ",forPrintStruct.count[t],"\n")
					t += 1
				} else{
					break
				}
			}

		}
	}

}
