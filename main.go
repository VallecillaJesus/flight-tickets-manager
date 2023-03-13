package main

import (
	"challenge/internal/tickets"
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Prompt{
		Label: "Path to the tickets csv file",
		Default: "tickets.csv",
	}
	csvFilePath, err := prompt.Run()

	if err != nil {
		panic(err)
	}
	
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

	switch itemIndex {
	case 0: 

		prompt = promptui.Prompt{
			Label: "Destination",
		}
		destination, err := prompt.Run()

		if err != nil {
			panic(err)
		}

		fmt.Println(t.GetTicketsAmountByDestination(destination))
	case 1: 
		prompt = promptui.Prompt{
			Label: "Start time",
		}
		st, err := prompt.Run()

		if err != nil {
			panic(err)
		}
		
		prompt = promptui.Prompt{
			Label: "End time",
		}
		et, err := prompt.Run()
		
		if err != nil {
			panic(err)
		}

		startTime := tickets.ParseToFlightTime(st)
		endTime := tickets.ParseToFlightTime(et)

		fmt.Println(t.GetTicketsAmountByTimeRange(startTime, endTime))
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

		switch itemIndex{
			case 0: fmt.Println(t.GetTicketsAmountByPeriod(tickets.EarlyMorning))
			case 1: fmt.Println(t.GetTicketsAmountByPeriod(tickets.Morning))
			case 2: fmt.Println(t.GetTicketsAmountByPeriod(tickets.Afternoon))
			case 3: fmt.Println(t.GetTicketsAmountByPeriod(tickets.Evening))
		}

	case 3: 
		prompt = promptui.Prompt{
			Label: "Destination",
		}
		destination, err := prompt.Run()

		if err != nil {
			panic(err)
		}

		prompt = promptui.Prompt{
			Label: "Start time",
		}
		st, err := prompt.Run()

		if err != nil {
			panic(err)
		}

		prompt = promptui.Prompt{
			Label: "End time",
		}
		et, err := prompt.Run()

		if err != nil {
			panic(err)
		}

		startTime, err := time.Parse("15:04",st)

		if err != nil {
			panic(err)
		}

		endTime, err := time.Parse("15:04",et)

		if err != nil {
			panic(err)
		}

		fmt.Println(t.GetTicketsPercentageByDestinationAndTimeRange(destination, startTime, endTime))

	case 4:
		fmt.Println(t.GetTicketsAverageByPeriods())
	}
}
