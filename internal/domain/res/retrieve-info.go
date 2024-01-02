package res

import (
	"time"
)

type RetrieveUserInformationResponse struct {
	Expires time.Time    `json:"expires"`
	Scopes  []string     `json:"scopes"`
	User    UserResponse `json:"user"`
}
