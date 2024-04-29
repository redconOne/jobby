package main

import (
	"fmt"
	"os"
	"strconv"

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

	return len(rows) - 1, nil
}

func AddJob(f *excelize.File, job Job) error {
	rows, err := f.GetRows("MyJobHunt")
	if err != nil {
		return err
	}
	newRow := len(rows) + 1
	headers := []string{
		"Company",
		"Website",
		"Role",
		"Description",
		"Skills",
		"Contacts",
		"Location",
		"Status",
		"Date",
		"Notes",
	}

	for _, key := range headers {
		switch key {
		case "Company":
			f.SetCellValue("MyJobHunt", "A"+strconv.Itoa(newRow), job.Company)
		case "Website":
			f.SetCellValue("MyJobHunt", "B"+strconv.Itoa(newRow), job.Website)
		case "Role":
			f.SetCellValue("MyJobHunt", "C"+strconv.Itoa(newRow), job.Role)
		case "Description":
			f.SetCellValue("MyJobHunt", "D"+strconv.Itoa(newRow), job.Description)
		case "Skills":
			fmt.Println(job.Skills)
			f.SetCellValue("MyJobHunt", "E"+strconv.Itoa(newRow), job.Skills)
		case "Contacts":
			f.SetCellValue("MyJobHunt", "F"+strconv.Itoa(newRow), job.Contacts)
		case "Location":
			f.SetCellValue("MyJobHunt", "G"+strconv.Itoa(newRow), job.Location)
		case "Status":
			f.SetCellValue("MyJobHunt", "H"+strconv.Itoa(newRow), job.Status)
		case "Date":
			f.SetCellValue("MyJobHunt", "I"+strconv.Itoa(newRow), job.Date)
		case "Notes":
			f.SetCellValue("MyJobHunt", "J"+strconv.Itoa(newRow), job.Notes)
		}
	}

	err = f.SaveAs("MyJobHunt.xlsx")
	if err != nil {
		return err
	}
	fmt.Println("Job application was successfully added!")
	return nil
}

// TODO: EditJob function
// TODO: GetStats function
