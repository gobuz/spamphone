package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func (rest *API) SpamPhones(r *gin.Context) {
	r.JSON(http.StatusOK, rest.PhoneService.GetSpamPhones())
}

func (rest *API) SpamPhone(r *gin.Context) {
	phoneNumber := r.Param("phone_number")
	if len(phoneNumber) > 0 {
		result := rest.PhoneService.GetSpamPhone(phoneNumber)
		if len(result.ID) > 0{
			r.JSON(http.StatusOK, result)
			return
		}
	}
	r.JSON(http.StatusNotFound, nil)
}