// Package tickets is a simple package in charge of managing
// and defining all the queries related to flight tickets, besides
// the definition of needed structs and types.
package tickets

import (
	"strings"
	"time"
	"log"
	"os"
	"strconv"
)

// ticket represents every flight ticket found in the external
// csv file that match with the struct model attributes order.
//
// Good csv file row content positions order example:
//
//	content: 1,Steve Musk,stevemusk@etsy.com,Colombia,20:44,550
//	positions:
//		[0]id = 1
//		[1]name = Steve Musk
//		[2]email = stevemusk@etsy.com
//		[3]destination = Colombia
//		[4]flightTime = 20:44
//		[5]price = 550
type ticket struct {
	id          string    // id is the flight ticket id.
	name        string    // name is the ticket passanger name.
	email       string    // email is the ticket passanger email.
	destination string    // destination is the flight ticket destination.
	flightTime  time.Time // flightTime is the time of the flight.
	price       float64   // price is the ticket price.
}

type Period [2]time.Time

// Tickets represents an slice containing all ticket structs.
// This is use to manipulate and query tickets struct data.
type Tickets []ticket

const flightTimeLayout = "15:04"

// ParseToFlightTime parses the given time string to the default
// flight accepted time layout, this is used to parse a time in 
// string to a Time instance.
// 
// The given time string is parsed using "15:04" time layout.
// Good time string should be "hours:minutes", example "10:30"
func ParseToFlightTime(t string) time.Time {
	parsedTime, err := time.Parse(flightTimeLayout, strings.TrimSpace(t))
	if err != nil {
		panic(err)
	}
	return parsedTime
}

var (
	EarlyMorning 	= Period{ParseToFlightTime("00:00"), ParseToFlightTime("06:59")}
	Morning 		= Period{ParseToFlightTime("07:00"), ParseToFlightTime("12:59")}
	Afternoon 		= Period{ParseToFlightTime("13:00"), ParseToFlightTime("19:59")}
	Evening 		= Period{ParseToFlightTime("20:00"), ParseToFlightTime("23:59")}
)


// GetTicketsAmountByDestination counts and returns the amount
// of flight tickets going to an specific destination.
func (t Tickets) GetTicketsAmountByDestination(destination string) int {
	var amount int
	for _, ticket := range t {
		if strings.EqualFold(destination, ticket.destination) {
			amount++
		}
	}
	return amount
}

func (t Tickets) GetTicketsAverageByPeriods() int {
	return	(t.GetTicketsAmountByPeriod(EarlyMorning) +
			t.GetTicketsAmountByPeriod(Morning) +
			t.GetTicketsAmountByPeriod(Afternoon) +
			t.GetTicketsAmountByPeriod(Evening) ) / 4
}

// GetTicketsAmountByTimeRange counts and returns the number of
// tickets which flightTime attribute is between the range of
// the given start time and end time.
func (t Tickets) GetTicketsAmountByTimeRange(startTime time.Time, endTime time.Time) int {
	var amount int
	for _, ticket := range t {
		if ticket.flightTime.After(startTime) && ticket.flightTime.Before(endTime) {
			amount++
		}
	}
	return amount
}

// GetTicketsAverageByDestinationAndTimeRange calculates the percentage of
// tickets going to an specific destination in a time range.
// 
// It returns the amount of tickets in a time range multiplied by 100 and
// divided by the amount of tickets going to a destination in the same time range.
// 
// 		percentage = ticketsAmountByDestinationAndTimeRange * 100 /
// 				  ticketsAmountByTimeRange   

func (t Tickets) GetTicketsPercentageByDestinationAndTimeRange(destination string, startTime time.Time, endTime time.Time) float64 {
	ticketsAmountByDestination := t.GetTicketsAmountByDestination(destination)
	ticketsAmountByTimeRange := t.GetTicketsAmountByTimeRange(startTime, endTime)
	return float64(ticketsAmountByDestination * 100) / float64(ticketsAmountByTimeRange)
}

func (t Tickets) GetTicketsAmountByPeriod(p Period) int {
	return t.GetTicketsAmountByTimeRange(p[0], p[1])
}

// ReadTickets reads the specified csv file path and transform
// each of the rows.
// Good csv file row content positions order example:
//
//	content: 1,Steve Musk,stevemusk@etsy.com,Colombia,20:44,550
//	positions:
//		[0]id = 1
//		[1]name = Steve Musk
//		[2]email = stevemusk@etsy.com
//		[3]destination = Colombia
//		[4]flightTime = 20:44
//		[5]price = 550
func ReadTickets(path string) (Tickets, error) {
	rawData, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var tickets Tickets

	data := strings.Split(strings.TrimSpace(string(rawData)), "\n")

	for _, row := range data {

		content := strings.Split(row, ",")	

		parsedFlightPrice, err := strconv.ParseFloat(content[5], 64)

		if err != nil {
			log.Fatal(err)
		}

		tickets = append(tickets, ticket{ 
			id: content[0],
			name: content[1],
			email: content[2],
			destination: content[3],
			flightTime: ParseToFlightTime(content[4]),
			price: parsedFlightPrice,
		})
	}
	return tickets, nil
}