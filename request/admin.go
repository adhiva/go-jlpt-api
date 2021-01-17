package request

type Admin struct {
	FirstName              string `json:"first_name" validate:"required"`
	LastName               string `json:"last_name" validate:"required"`
	Email                  string `json:"email" validate:"required,email"`
	Password               string `json:"password" validate:"required"`
	PhoneNumberCountryCode string `json:"phone_number_country_code" validate:"required"`
	PhoneNumber            string `json:"phone_number" validate:"required"`
	Country                string `json:"country" validate:"required"`
}
