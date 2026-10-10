package validator

var enMessages = map[string]map[string]string{
	//Register
	"FullName": {
		"required": "Full name is required.",
	},

	"PhoneNumber": {
		"required": "Phone number cannot be empty.",
		"min":      "Phone number must be at least 12 characters.",
		"max":      "Phone number must be at most 20 characters.",
	},

	"Email": {
		"required": "Email cannot be empty.",
		"email":    "The entered email is not valid.",
		"min":      "Email must be at least 3 characters.",
		"max":      "Email must be at most 250 characters.",
	},

	"Password": {
		"required": "Password is required.",
		"min":      "Password must be at least 8 characters.",
	},

	"RepeatPassword": {
		"required": "Repeat password is required.",
	},

	//Login
	"EmailOrPhone": {
		"required": "Email or phone number is required.",
		"max":      "Email or phone number cannot exceed 254 characters",
	},

	//VerificationToken
	"VerificationToken": {
		"required": "verification token is required.",
	},

	"Code": {
		"required": "Code is required.",
	},

	"CurrentPassword": {
		"min": "Current password must be at least 8 characters.",
	},

	"NewPassword": {
		"min": "New password must be at least 8 characters.",
	},
}
