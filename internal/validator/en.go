package validator

var enMessages = map[string]map[string]string{
	//Register
	"FullName": {
		"required": "Full name is required.",
	},

	"PhoneNumber": {
		"required": "Phone number is required.",
	},

	"Email": {
		"required": "Email is required.",
		"email":    "Please enter a valid email address.",
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
