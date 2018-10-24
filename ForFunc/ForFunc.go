package ForFunc

import "fmt"

type ForPrint struct {
	namefi string
	count  int
}

var forFrintStr []ForPrint // Здесь храним данные для вывода

func AddStruct(count int, file string) int {

	var Test ForPrint

	Test.count = count
	Test.namefi = file
	forFrintStr = append(forFrintStr, Test)

	return 1
}

func SortStruct() {
	for i := 0; i < len(forFrintStr)-1; i += 1 {
		for j := i; j < len(forFrintStr); j += 1 {
			if forFrintStr[i].count < forFrintStr[j].count {
				forFrintStr[i].count, forFrintStr[j].count = forFrintStr[j].count, forFrintStr[i].count
				forFrintStr[i].namefi, forFrintStr[j].namefi = forFrintStr[j].namefi, forFrintStr[i].namefi
			}
		}
	}
	//fmt.Println(forFrintStr)
	for i := 0; i < len(forFrintStr); i += 1 {
		fmt.Println("-", forFrintStr[i].namefi, "; совпадений -", forFrintStr[i].count)
	}
}
