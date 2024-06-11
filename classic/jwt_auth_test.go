package classic_test

import (
	"testing"
	"time"

	jamf "github.com/DataDog/jamf-api-client-go/classic"
	"github.com/stretchr/testify/assert"
)

func TestJWTAuthExpiration(t *testing.T) {
	type expirationCase struct {
		msg             string
		exp             string
		shouldBeExpired bool
	}

	now := time.Now()

	cases := []expirationCase{
		{
			msg:             "token expires 30 minutes from now",
			exp:             now.Add(time.Minute * 30).Format(time.RFC3339),
			shouldBeExpired: false,
		},
		{
			msg:             "token expired 5 minutes ago",
			exp:             now.Add(time.Minute * -5).Format(time.RFC3339),
			shouldBeExpired: true,
		},
		{
			msg:             "token expires in 30 seconds",
			exp:             now.Add(time.Second * 30).Format(time.RFC3339),
			shouldBeExpired: false,
		},
		{
			msg:             "token expired 1 second ago",
			exp:             now.Add(time.Second * -1).Format(time.RFC3339),
			shouldBeExpired: true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.msg, func(t *testing.T) {
			t.Parallel() // marks each test case as capable of running in parallel with each other
			token := jamf.AuthToken{
				Token:   "test-token",
				Expires: c.exp,
			}

			expired, err := token.IsExpired()
			assert.Nil(t, err)
			assert.Equal(t, c.shouldBeExpired, expired)
		})

	}

}
