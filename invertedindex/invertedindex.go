package invertedindex

var vocabulary map[string]map[string]int // слово → имя файла → количество вхождений
var fileList map[string]map[string]int   // имя файла → слово → количество вхождений

type ForPrint struct {
	Namefi string
	Count  int
}

var ForPrintStr []ForPrint // Здесь храним данные для вывода

func AddStruct(count int, file string) int { // Добавили файл в вывод
	var Test ForPrint
	Test.Count = count
	Test.Namefi = file

	ForPrintStr = append(ForPrintStr, Test)
	return 1
}

func SortStruct() {
	for i := 0; i < len(ForPrintStr)-1; i += 1 {
		for j := i; j < len(ForPrintStr); j += 1 {
			if ForPrintStr[i].Count < ForPrintStr[j].Count {
				ForPrintStr[i].Count, ForPrintStr[j].Count = ForPrintStr[j].Count, ForPrintStr[i].Count
				ForPrintStr[i].Namefi, ForPrintStr[j].Namefi = ForPrintStr[j].Namefi, ForPrintStr[i].Namefi
			}
		}
	}
}

func AddMap() bool {
	vocabulary = make(map[string]map[string]int, 1)
	fileList = make(map[string]map[string]int, 1)
	return true
}

func AddNewFile(words []string, file string) bool {

	for _, word := range words {

		if word != " " {
			if _, ok := vocabulary[word]; ok {
				if _, ok1 := vocabulary[word][file]; ok1 {
					vocabulary[word][file] += 1
				} else {
					vocabulary[word][file] = 1
				}
			} else {
				tmp := make(map[string]int)
				tmp[file] = 1
				vocabulary[word] = tmp
			}
			if _, ok := fileList[file]; ok {
				if _, ok1 := fileList[file][word]; ok1 {
					fileList[file][word] += 1
				} else {
					fileList[file][word] = 1
				}
			} else {
				tmp := make(map[string]int)
				tmp[word] = 1
				fileList[file] = tmp
			}
		}

	}
	return true
}

func FileSearch(example []string) []ForPrint {

	inst := 0
	for word, tmp := range vocabulary { // Ищем первое слово в словаре
		// !!! Что делать с пустым запросом
		if word == example[0] { // Нашли файл с совпадением
			for file, count := range tmp { // проходим поочередно по всем файлам, в которых есть первое слово
				// ищем каждый файл во вспомогательной map
				countwrite := count

				for file1, _ := range fileList {
					flag := 1
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
						if flag == len(example) { // значит встретили все слова из запроса → можем выводить на экран
							flag = AddStruct(countwrite, file)
							inst += 1
						}
						break
					}
				}
			}
			break
		}
	}
	SortStruct()
	return ForPrintStr
}
