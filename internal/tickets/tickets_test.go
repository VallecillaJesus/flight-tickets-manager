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


