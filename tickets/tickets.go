// Package tickets is a simple package in charge of managing
// and defining all the queries related to flight tickets, besides 
// the definition of needed structs and types.
package tickets

import (
	"strings"
	"time"
)

// ticket represents every flight ticket found in the external
// csv file that match with the struct model attributes order.
//
// Good csv file row content positions order example:
//
// 		content: 1,Steve Musk,stevemusk@etsy.com,Colombia,20:44,550
// 		positions: 
// 			[0]id = 1
// 			[1]name = Steve Musk
// 			[2]email = stevemusk@etsy.com
// 			[3]destination = Colombia
//			[4]flightTime = 20:44
// 			[5]price = 550
type ticket struct {
	id string 				// id is the flight ticket id.
	name string 			// name is the ticket passanger name.
	email string 			// email is the ticket passanger email.
	destination string 		// destination is the flight ticket destination.
	flightTime time.Time	// flightTime is the time of the flight.
	price float64 			// price is the ticket price.
}

// Tickets represents an slice containing all ticket structs.
// This is use to manipulate and query tickets struct data.
type Tickets []ticket

// GetTicketsAmountByDestination counts and returns the amount 
// of flight tickets going to an specific destination.
func (t Tickets) GetTicketsAmountByDestination(destination string) int {
	var amount int
	for _, ticket := range t {
		if strings.ToLower(destination) == strings.ToLower(ticket.destination) {
			amount++
		}
	}
	return amount
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

func (t Tickets) GetTicketsAverageByDestination(destination string) float64 {
	var amount float64
	for _, ticket := range t {
		if strings.ToLower(destination) == strings.ToLower(ticket.destination) {
			amount++
		}
	}
	return amount / float64(len(t))
}
