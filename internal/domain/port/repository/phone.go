package repository

import "github.com/gobuz/publicspam/internal/domain/model/phone"

type Repository interface {
	GetAllSpamPhones() []phone.SpamPhone
	GetSpamPhone(phoneNumber string) phone.SpamPhone
}
