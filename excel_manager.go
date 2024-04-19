package main

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func InitializeOrLoadExcelFile(filename string) (*excelize.File, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		f := excelize.NewFile()
		sheet := "MyJobHunt"

		index, err := f.NewSheet(sheet)
		if err != nil {
			return nil, err
		}

		headers := []string{"Company", "Website", "Role", "Description", "Skills", "Contacts", "Location", "Status", "Date", "Notes"}

		for i, header := range headers {
			cell, _ := excelize.CoordinatesToCellName(i+1, 1)
			f.SetCellValue(sheet, cell, header)
		}

		f.SetActiveSheet(index)

		if err := f.SaveAs(filename); err != nil {
			return nil, err
		}

		fmt.Printf("New file created: %s\n", filename)
		return f, nil
	} else {
		f, err := excelize.OpenFile(filename)
		if err != nil {
			return nil, err
		}

		fmt.Printf("File loaded: %s\n", filename)
		return f, nil
	}
}

func CountJobs(f *excelize.File) (int, error) {
	rows, err := f.GetRows("MyJobHunt")
	if err != nil {
		return 0, err
	}

	numJobs := 0

	for _, row := range rows {
		if row[0] != "" {
			numJobs++
		}
	}

	return numJobs - 1, nil
}

// TODO: AddJob function
// TODO: EditJob function
// TODO: GetStats function
