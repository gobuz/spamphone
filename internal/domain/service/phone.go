package service

import (
	"github.com/gobuz/publicspam/internal/domain/model/phone"
	"github.com/gobuz/publicspam/internal/domain/port/repository"
	"github.com/gobuz/publicspam/internal/domain/port/service"
)

type phoneService struct {
	PhoneRepository repository.Repository
}

// InitPhoneService init service of phone
func InitPhoneService(r repository.Repository) service.Phone {
	return &phoneService{
		PhoneRepository: r,
	}
}

func (service *phoneService) GetSpamPhones() []phone.SpamPhone {
	return service.PhoneRepository.GetAllSpamPhones()
}

func (service *phoneService) GetSpamPhone(phoneNumber string) phone.SpamPhone {
	return service.PhoneRepository.GetSpamPhone(phoneNumber)
}