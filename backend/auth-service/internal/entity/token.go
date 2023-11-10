package entity

import (
	"encoding/json"
	"time"
)

type Token struct {
	Access           string
	AccessExpiresIn  time.Duration
	Refresh          string
	RefreshExpiresIn time.Duration
}

func (t Token) MarshalJSON() ([]byte, error) {
	var tmp struct {
		Access    string `json:"access,omitempty"`
		ExpiresIn int    `json:"expires_in,omitempty"`
	}

	tmp.Access = t.Access
	tmp.ExpiresIn = int(t.AccessExpiresIn.Seconds())
	return json.Marshal(&tmp)
}
