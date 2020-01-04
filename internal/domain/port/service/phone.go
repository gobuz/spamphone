package service

import "github.com/gobuz/publicspam/internal/domain/model/phone"

type Phone interface {
	GetSpamPhones() []phone.SpamPhone
	GetSpamPhone(phoneNumber string) phone.SpamPhone
}
