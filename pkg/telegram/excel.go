package telegram

import (
	"github.com/tealeg/xlsx"
	"log"
	"strconv"
)

func searchGroup(nameGroup string) bool {
	wb, err := xlsx.OpenFile(fileIiVT)
	if err != nil {
		log.Fatal(err)
		return false
	}
	for i := 1; i <= 14; i++ {
		sh, ok := wb.Sheet[strconv.Itoa(i)]
		if ok {

			// F2
			theCell := sh.Cell(1, 5)
			if nameGroup == theCell.String() {
				return true
			}

			// K2
			theCell = sh.Cell(1, 10)
			if nameGroup == theCell.String() {
				return true
			}
		}
	}
	return false
}
