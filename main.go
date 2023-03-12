package main

import (
	"fmt"
	"time"
	"challenge/tickets"
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
			"Average of tickets by destination", 
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

		startTime, err := time.Parse("15:04",st)

		if err != nil {
			panic(err)
		}

		endTime, err := time.Parse("15:04",et)

		if err != nil {
			panic(err)
		}

		fmt.Println(t.GetTicketsAmountByTimeRange(startTime, endTime))
	case 2: 
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
	}
}
