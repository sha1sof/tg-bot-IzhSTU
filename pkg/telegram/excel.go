package telegram

import (
	"github.com/tealeg/xlsx"
	"log"
	"strconv"
)

// return found(bool), row(int), col(int), sheet
func searchGroup(nameGroup string) (bool, int, int, string) {
	wb, err := xlsx.OpenFile(fileIiVT)
	if err != nil {
		log.Fatal(err)
		return false, 0, 0, ""
	}
	for i := 1; i <= 14; i++ {
		sh, ok := wb.Sheet[strconv.Itoa(i)]
		if ok {
			// F2
			theCell := sh.Cell(1, 5)
			if nameGroup == theCell.String() {
				return true, 1, 5, sh.Name
			}

			// K2
			theCell = sh.Cell(1, 10)
			if nameGroup == theCell.String() {
				return true, 1, 10, sh.Name
			}
		}
	}
	return false, 0, 0, ""
}
