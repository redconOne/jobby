package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type Job struct {
	Date        string
	Company     string
	Website     string
	Role        string
	Description string
	Location    string
	Status      Status
	Notes       string
	Skills      []string
	Contacts    []Contact
}

type Contact struct {
	Name     string // Contact's name
	Platform string // Platform Contact was reached out to
	Notes    string // Additional information about contact
}

type Status string

const (
	Applied      Status = "applied"
	Stale        Status = "stale"
	Interviewing Status = "interviewing"
	Offer        Status = "offer received"
)

func main() {
	f, err := InitializeOrLoadExcelFile("MyJobHunt.xlsx")
	if err != nil {
		fmt.Printf("Failed to initialize or load Excel file: %v\n", err)
		return
	}
	defer f.Close()

	var (
		operation       string
		job             Job
		jobSkills       string
		jobContact      string
		contactPlatform string
	)

	numJobs, err := CountJobs(f)
	if err != nil {
		fmt.Printf("Failed to get number of jobs: %v\n", err)
	}

	// Should we run in accessible mode?
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

	mainMenu := huh.NewForm(
		huh.NewGroup(huh.NewNote().
			Title("Welcome").
			Description("Welcome to Jobby™.")),

		// Choose an option.
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("Add a new job application", "Edit existing application", "View statistics", "Exit")...).
				Title("Main Menu").
				Description("Please choose one option").
				Validate(func(t string) error {
					if t == "Edit existing application" && numJobs < 1 {
						return fmt.Errorf("you have no job applications saved")
					}
					return nil
				}).
				Value(&operation),
		),
	).WithAccessible(accessible)

	err = mainMenu.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	prepareOperation := func() {
		time.Sleep(2 * time.Second)
	}

	_ = spinner.New().Title("One moment please...").Accessible(accessible).Action(prepareOperation).Run()

	// TODO: Handle main menu options (Add, Edit, Stats, Exit)

	switch operation {
	case "Add a new job application":
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Company name:").
					Value(&job.Company),
				huh.NewInput().
					Title("Company website:").
					Value(&job.Website),
			),
			huh.NewGroup(
				huh.NewInput().
					Title("Role:").
					Value(&job.Role),
				huh.NewText().
					Title("Job Description:").
					Value(&job.Description),
				huh.NewInput().
					Title("Required Skills:").
					Description("Separated by , EG[Go, Python, Rust, etc]").
					Value(&jobSkills),
				huh.NewInput().
					Title("Location:").
					Description("Please use remote, local, or city/state").
					Value(&job.Location),
			),
			huh.NewGroup(
				huh.NewInput().
					Title("Contact:").
					Value(&jobContact),
				huh.NewInput().
					Title("Platform:").
					Description("Where did you meet/contact this person?").
					Value(&contactPlatform),
			),
			huh.NewGroup(
				huh.NewText().
					Title("Additional Notes:").
					Value(&job.Notes),
			),
		).WithAccessible(accessible)

		err = form.Run()
		if err != nil {
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}

		job.Skills = strings.Split(jobSkills, ", ")
		job.Contacts = []Contact{{Name: jobContact, Platform: contactPlatform}}
		year, month, day := time.Now().Date()
		job.Date = fmt.Sprintf("%d-%02d-%02d", year, month, day)
		job.Status = Applied

		AddJob(f, job)
	}
}
