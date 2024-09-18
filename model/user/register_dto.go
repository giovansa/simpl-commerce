package user

import (
	"regexp"
	repository "simpl-commerce/repository/user"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserRequest struct {
	Phone    string `json:"phone" validate:"required,min=10,max=13,phone_prefix=+62"`
	Email    string `json:"email" validate:"required,min=10,max=30,email"`
	Name     string `json:"name" validate:"required,min=3,max=60"`
	Password string `json:"password" validate:"required,min=6,max=64,password"`
}

func (r *RegisterUserRequest) Validate() error {
	validate := validator.New()
	registerCustomValidators(validate)
	return validate.Struct(r)
}

func (r *RegisterUserRequest) ToDAO() (repository.RegisterUser, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return repository.RegisterUser{}, err
	}
	return repository.RegisterUser{
		ID:       uuid.New().String(),
		Phone:    r.Phone,
		Name:     r.Name,
		Password: string(hashedPassword),
	}, nil
}

func validateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, email)
	return match
}

func validatePhonePrefix(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	return strings.HasPrefix(phone, "+62")
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	var (
		hasSpecialChar bool
		hasCapital     bool
		hasNumeric     bool
	)

	for _, char := range password {
		switch {
		case unicode.IsDigit(char):
			hasNumeric = true
		case unicode.IsUpper(char):
			hasCapital = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
	}

	return hasSpecialChar && hasCapital && hasNumeric
}

func registerCustomValidators(validate *validator.Validate) {
	err := validate.RegisterValidation("phone_prefix", validatePhonePrefix)
	if err != nil {
		return
	}
	err = validate.RegisterValidation("password", validatePassword)
	if err != nil {
		return
	}
	err = validate.RegisterValidation("email", validateEmail)
	if err != nil {
		return
	}
}
