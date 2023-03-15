package main

import (
	"challenge/internal/tickets"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
)

var (
	errEmptyPrompt = errors.New("input can not be empty")
	errInvalidTimeString = errors.New("time format is not valid, valid format is '11:00'")
	errInvalidCSVFile = errors.New("not valid csv file found")
)

func validateEmptyPrompt(s string) error {
	if len(s) == 0 {
		return errEmptyPrompt
	}
	return nil
}

func validateTimeString(s string) error {
	_, err := time.Parse(tickets.FlightTimeLayout, s)
	if err != nil {
		return errInvalidTimeString
	}
	return nil
}

func validateCSVFile(s string) error {
	if  _, err := os.Open(s); err != nil {
		return errInvalidCSVFile
	}

	format := strings.Split(s, ".")[1]
	if format != "csv" {
		return errInvalidCSVFile
	}

	return nil
}

func main() {

	var result string

	prompt := promptui.Prompt{
		Label: "Path to the tickets csv file",
		Default: "tickets.csv",
		Validate: validateCSVFile,
	}

	csvFilePath, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	prompt = promptui.Prompt{}

	s := spinner.New(spinner.CharSets[14], 50 * time.Millisecond)
	s.Color("yellow", "bold")
	s.FinalMSG = "✔ File read successfully\n"
	s.Suffix = " Reading file..."
	s.Start()

	t, err := tickets.ReadTickets(csvFilePath)
	time.Sleep(1 * time.Second)

	
	if err != nil {
		s.FinalMSG = "❌ Fail while reading the file\n"
		s.Stop()
		panic(err)
	}

	s.Stop()

	// Empty line
	fmt.Println()

	selector := promptui.Select{
		Label: "Select the information you need",
		Items: []string{
			"Amount of tickets by destination", 
			"Amount of tickets by time range",
			"Amount of tickets by period",
			"Percentage of tickets by destination and time range",
			"Average of tickets by periods",
		},
	}

	itemIndex, _, err := selector.Run()

	if err != nil {
		panic(err)	
	}

	fmt.Println("\nFill the prompts, use `↩ enter` key to display more information.")

	// Handle selected selector Item.
	switch itemIndex {
	// Amount of tickets by destination.
	case 0: 
		prompt := promptui.Prompt{
			Label: "Destination",
			Validate: validateEmptyPrompt,
		}
		destination, err := prompt.Run()

		if err != nil {
			panic(err)
		}
		
		result = fmt.Sprintf("The tickets amount is %d ", t.GetTicketsAmountByDestination(destination))
	
	// Amount of tickets by time range.
	case 1: 
		prompt = promptui.Prompt{
			Label: "Start time",
			Validate: validateTimeString,
		}
		st, err := prompt.Run()

		if err != nil {
			panic(err)
		}
		
		prompt = promptui.Prompt{
			Label: "End time",
			Validate: validateTimeString,
		}
		et, err := prompt.Run()
		
		if err != nil {
			panic(err)
		}

		result = fmt.Sprintf("The tickets amount is %d ", t.GetTicketsAmountByTimeRange(
			tickets.ParseToFlightTime(st), 
			tickets.ParseToFlightTime(et),
		))
	
	// Amount of tickets by period.
	case 2:
		selector = promptui.Select{
			Label: "Select the period",
			Items: []string{
				"Early morning", 
				"Morning",
				"Afternoon",
				"Evening", 
			},
		}

		itemIndex, _, err := selector.Run()

		if err != nil {
			panic(err)
		}

		msg := "The tickets amount is %d "
		
		// Handle selected `Period`.
		switch itemIndex{
			// Early morning.
			case 0: result = fmt.Sprintf(msg, t.GetTicketsAmountByPeriod(tickets.EarlyMorning))
			// Morning.
			case 1: result = fmt.Sprintf(msg, t.GetTicketsAmountByPeriod(tickets.Morning))
			// Afternoon.
			case 2: result = fmt.Sprintf(msg, t.GetTicketsAmountByPeriod(tickets.Afternoon))
			// Evening.
			case 3: result = fmt.Sprintf(msg, t.GetTicketsAmountByPeriod(tickets.Evening))
		}

	// Percentage of tickets by destination and time range.
	case 3: 
		prompt = promptui.Prompt{
			Label: "Destination",
			Validate: validateEmptyPrompt,
		}
		destination, err := prompt.Run()

		if err != nil {
			panic(err)
		}

		prompt = promptui.Prompt{
			Label: "Start time",
			Validate: validateTimeString,
		}
		st, err := prompt.Run()

		if err != nil {
			panic(err)
		}

		prompt = promptui.Prompt{
			Label: "End time",
			Validate: validateTimeString,
		}
		et, err := prompt.Run()

		if err != nil {
			panic(err)
		}

		result = fmt.Sprintf("The percentage of tickets is %.2f%%", t.GetTicketsPercentageByDestinationAndTimeRange(
			destination,
			tickets.ParseToFlightTime(st), 
			tickets.ParseToFlightTime(et),
		))

	// Average of tickets by periods.
	case 4:
		result = fmt.Sprintf("The average of tickets is %.2f",t.GetTicketsAverageByPeriods())
	}

	fmt.Println("\n⭐️ " + result + "\n")
}
