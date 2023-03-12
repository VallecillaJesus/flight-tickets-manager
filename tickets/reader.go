package tickets

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)


// ReadTickets reads the specified csv file path and transform the each of the rows in
func ReadTickets (path string) (Tickets, error) {
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