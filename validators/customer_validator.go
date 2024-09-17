package validators

import (
	"fmt"
	"time"

	"github.com/olmandaniel/customers.api/models"
	"github.com/olmandaniel/customers.api/utils"
)

func ValidatorCustomer(customer models.Customer) (bool, error) {
	layout := "2006-01-02"
	date, err := time.Parse(layout, customer.Birthdate)
	if err != nil {
		return false, err
	}
	fmt.Println(customer.Birthdate)
	duration := time.Since(date)
	if utils.DurationToYears(duration) < 18 {
		return false, fmt.Errorf("solo se pueden registrar personas mayores de 18 años")
	}

	if !utils.ValidateDNI(customer.Document) {
		return false, fmt.Errorf("el dni no es válido")
	}

	if !utils.ValidatePhone(customer.Phone) {
		return false, fmt.Errorf("el celular no es válido")
	}

	if !utils.ValidateEmail(customer.Email) {
		return false, fmt.Errorf("el email no es válido")
	}

	return true, nil
}
