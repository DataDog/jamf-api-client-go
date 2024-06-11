package classic

import (
	"time"
)

type AuthToken struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}

func (t *AuthToken) IsExpired() (bool, error) {
	if t.Expires == "" {
		return true, nil
	}
	expiration, err := time.Parse(time.RFC3339, t.Expires)
	if err != nil {
		return true, err
	}
	return expiration.Before(time.Now()), nil
}
