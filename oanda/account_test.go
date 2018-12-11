package oanda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApi_GetAccount(t *testing.T) {
	 api := StubResponse("../fixture/account.json")

	res, _ := api.GetAccount()
	assert.Equal(t, res.Account.ID, "your own id")
}
