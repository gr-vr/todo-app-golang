package todo

import "time"

type User struct {
	ID              int       `json:"-" db:"id"`
	Created         time.Time `json:"created"`
	Updated         time.Time `json:"updated"`
	AccountType     string    `json:"account_type"`
	Email           string    `json:"email" binding:"required"`
	Verified        bool      `json:"verified"`
	Password        string    `json:"password" binding:"required"`
	Name            string    `json:"name" binding:"required"`
	LastName        string    `json:"last_name" binding:"required"`
	OrganizationID  int64     `json:"organization_id"`
	Timezone        string    `json:"timezone"`
	Language        string    `json:"language"`
	PaymentType     string    `json:"payment_type"`
	PaymentID       string    `json:"payment_id"`
	ValidUntil      time.Time `json:"valid_until"`
	SlidesAvailable int64     `json:"slides_available"`
}
