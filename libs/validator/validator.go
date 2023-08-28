package error_handler

import (
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CustomErrorHandler(err error) fiber.Map {
	msg := fiber.Map{}

	for _, e := range err.(validator.ValidationErrors) {
		field := toSnakeCase(e.Field()) // Convert field name to lowercase
		switch e.Tag() {
		case "required":
			msg[field] = "mohon isi " + field
		case "min":
			msg[field] = "panjang minimum " + field + " adalah " + e.Param()
		case "email":
			msg[field] = "mohon masukkan email yang benar"
		default:
			msg[field] = "invalid " + field
		}
	}
	return msg

}

func SqlErrorHandler(err string) fiber.Map {
	if err == "pq: duplicate key value violates unique constraint \"users_email_key\"" {

		return fiber.Map{
			"email": "email sudah terdaftar",
		}
	}
	return fiber.Map{
		"error":err,
	}
}

func toSnakeCase(s string) string {
	var output []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				output = append(output, '_')
			}
			output = append(output, unicode.ToLower(r))
		} else {
			output = append(output, r)
		}
	}
	return string(output)
}
