package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheet := "MyJobHunt"

	index, err := f.NewSheet(sheet)
	if err != nil {
		fmt.Println(err)
		return
	}

	f.SetCellValue(sheet, "A1", "Company")
	f.SetCellValue(sheet, "B1", "Website")
	f.SetCellValue(sheet, "C1", "Role/Position")
	f.SetCellValue(sheet, "D1", "Description")
	f.SetCellValue(sheet, "E1", "Required Skills")
	f.SetCellValue(sheet, "F1", "Contacts")
	f.SetCellValue(sheet, "G1", "Location")

	f.SetActiveSheet(index)

	if err := f.SaveAs("MyJobHunt.xlsx"); err != nil {
		fmt.Println(err)
	}
}
