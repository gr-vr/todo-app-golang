package todo

import "time"

type Organization struct {
	ID               int64     `json:"id"`
	Created          time.Time `json:"created"`
	Updated          time.Time `json:"updated"`
	AccountType      string    `json:"account_type"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	OrganizationName string    `json:"organization_name"`
	Timezone         string    `json:"timezone"`
	Language         string    `json:"language"`
	PaymentType      string    `json:"payment_type"`
	PaymentID        string    `json:"payment_id"`
	ValidUntil       time.Time `json:"valid_until"`
	UsersAvailable   int64     `json:"users_available"`
}
