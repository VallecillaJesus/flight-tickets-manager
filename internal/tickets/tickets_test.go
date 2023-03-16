package tickets

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReadTickets(t *testing.T) {
	tickets, err := ReadTickets("mock.csv")
	assert.Nil(t, err)
	assert.NotNil(t, tickets)
	assert.EqualValues(t, 1000, len(tickets))
}


func TestParseToFlightTimeInvalidFormat(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r, "The function did not panic")
	}()
	// Invalid time string format
	ParseToFlightTime("0000")
}

func TestParseToFlightTime(t *testing.T) {
	parsedTime := ParseToFlightTime("00:00")
	assert.NotEmpty(t, parsedTime, "The parsed time returned is empty")

	parsedTime = ParseToFlightTime("23:59")
	assert.NotEmpty(t, parsedTime, "The parsed time returned is empty")
}

func TestGetTicketsAmountByDestination(t *testing.T) {
	tickets, _ := ReadTickets("mock.csv")
	amount := tickets.GetTicketsAmountByDestination("Colombia")
	assert.EqualValues(t, 18, amount)

	amount = tickets.GetTicketsAmountByDestination("Argentina")
	assert.EqualValues(t, 15, amount)

	amount = tickets.GetTicketsAmountByDestination("Neverland")
	assert.EqualValues(t, 0, amount)
}
