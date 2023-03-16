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
