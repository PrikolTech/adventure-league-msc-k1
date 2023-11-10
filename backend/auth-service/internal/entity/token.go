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
	UserID           string
}

func (t Token) MarshalJSON() ([]byte, error) {
	var tmp struct {
		Access    string `json:"access,omitempty"`
		ExpiresIn int    `json:"expires_in,omitempty"`
		UserID    string `json:"user_id,omitempty"`
	}

	tmp.Access = t.Access
	tmp.ExpiresIn = int(t.AccessExpiresIn.Seconds())
	tmp.UserID = t.UserID
	return json.Marshal(&tmp)
}
