package phone

import (
	"time"
)

type SpamPhone struct {
	ID string `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
	Note string `json:"note""`
	Type string `json:"type"`
	ReportedTimes int `json:"reportedTimes"`
	Reports []Report `json:"reports"`
}

type Report struct {
	UserId string `json:"userID"`
	Note string `json:"note"`
	Type string `json:"type"`
	CreatedDate time.Time `json:"createdDate"'`
}