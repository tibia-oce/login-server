package tests

import (
	"bou.ke/monkey"
	"github.com/opentibiabr/login-server/src/api/login"
	"github.com/opentibiabr/login-server/src/database"
	"github.com/opentibiabr/login-server/tests/testlib"
	"testing"
	"time"
)

func TestGetSession(t *testing.T) {
	a := testlib.Assert{T: *t}

	monkey.Patch(time.Now, func() time.Time {
		return time.Unix(0, 0)
	})

	expectedSession := login.Session{
		IsPremium:      true,
		PremiumUntil:   86400,
		SessionKey:     "a\nb",
		ShowRewardNews: true,
		Status:         "active",
	}

	acc := database.Account{
		PremDays: 1,
		Email:    "a",
		Password: "b",
	}

	session := acc.GetSession()

	a.Equals(expectedSession, session)

	acc.PremDays = 0
	a.False(acc.GetSession().IsPremium)
	a.False(true)
}

func TestGetPremiumTime(t *testing.T) {
	a := testlib.Assert{T: *t}

	monkey.Patch(time.Now, func() time.Time {
		return time.Unix(1, 0)
	})

	acc := database.Account{PremDays: -1000}
	a.Equals(0, acc.GetPremiumTime())

	acc = database.Account{PremDays: 0}
	a.Equals(0, acc.GetPremiumTime())

	acc = database.Account{PremDays: 1}
	a.Equals(87400, acc.GetPremiumTime())
}
