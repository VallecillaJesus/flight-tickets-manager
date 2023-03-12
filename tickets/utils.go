package tickets

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// ParseToFlightTime parses the given time string to the default
// flight accepted time layout, this is used to parse a time in 
// string to a Time instance.
// 
// The given time string is parsed using "15:04" time layout.
// Good time string should be "hours:minutes", example "10:30"
func ParseToFlightTime(t string) time.Time {
	parsedTime, err := time.Parse("15:04", t)
	if err != nil {
		panic(err)
	}
	return parsedTime
}

// ReadTickets reads the specified csv file path and transform the each of the rows in
func ReadTickets(path string) (Tickets, error) {
	rawData, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var tickets Tickets

	data := strings.Split(strings.TrimSpace(string(rawData)), "\n")

	for _, row := range data {

		content := strings.Split(row, ",")	

		parsedFlightTime, err := time.Parse("15:04", content[4])

		if err != nil {
			log.Fatal(err)
		}

		parsedFlightPrice, err := strconv.ParseFloat(content[5], 64)

		if err != nil {
			log.Fatal(err)
		}

		tickets = append(tickets, ticket{ 
			id: content[0],
			name: content[1],
			email: content[2],
			destination: content[3],
			flightTime: parsedFlightTime,
			price: parsedFlightPrice,
		})
	}
	return tickets, nil
}