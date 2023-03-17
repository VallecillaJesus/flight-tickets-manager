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

func TestGetTicketsAverageByPeriods(t *testing.T) {
	tickets, _ := ReadTickets("mock.csv")

	average := tickets.GetTicketsAverageByPeriods()
	assert.EqualValues(t, 250.0, average)
}

func TestGetTicketsAmountByPeriod(t *testing.T) {
	tickets, _ := ReadTickets("mock.csv")

	average := tickets.GetTicketsAmountByPeriod(EarlyMorning)
	assert.EqualValues(t, 303, average)


	average = tickets.GetTicketsAmountByPeriod(Morning)
	assert.EqualValues(t, 255, average)


	average = tickets.GetTicketsAmountByPeriod(Afternoon)
	assert.EqualValues(t, 289, average)


	average = tickets.GetTicketsAmountByPeriod(Evening)
	assert.EqualValues(t, 151, average)
}

func TestGetTicketsAmountByTimeRange(t *testing.T) {
	tickets, _ := ReadTickets("mock.csv")

	amount := tickets.GetTicketsAmountByTimeRange(ParseToFlightTime("10:00"), ParseToFlightTime("12:00"))
	assert.EqualValues(t, 81, amount)

}

func TestGetTicketsPercentageByDestinationAndTimeRange(t *testing.T) {
	tickets, _ := ReadTickets("mock.csv")
	amount := tickets.GetTicketsPercentageByDestinationAndTimeRange(
		"Colombia",
		ParseToFlightTime("00:00"),
		ParseToFlightTime("23:00"),
	)
	assert.EqualValues(t, 1.8769551616266944, amount)
}