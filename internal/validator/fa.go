package validator

var faMessages = map[string]map[string]string{
	"FullName": {
		"required": "نام و نام خانوادگی نباید خالی باشد.",
	},

	"PhoneNumber": {
		"required": "شماره تلفن نباید خالی باشد.",
		"min":      "شماره تلفن باید حداقل 12 کارکتر باشد",
		"max":      "شماره تلفن باید حداکثر 20 کارکتر باشد",
	},

	"Email": {
		"required": "ایمیل نباید خالی باشد.",
		"email":    "ایمیل وارد شده معتبر نیست.",
		"min":      "ایمیل باید حداقل 3 کارکتر باشد",
		"max":      "ایمیل باید حداکثر 250 کارکتر باشد",
	},

	"Password": {
		"required": "رمز عبور نباید خالی باشد.",
		"min":      "رمز عبور باید حداقل ۸ کاراکتر باشد.",
	},

	"RepeatPassword": {
		"required": "تکرار رمز عبور نباید خالی باشد.",
	},

	//Login
	"EmailOrPhone": {
		"required": "ایمیل یا تلفن نباید خالی باشد",
		"max":      "ایمیل یا شماره تلفن نباید بیشتر از ۲۵۴ کاراکتر باشد",
	},

	//VerificationToken
	"VerificationToken": {
		"required": "توکن تأیید الزامی است.",
	},

	"Code": {
		"required": "کد تأیید الزامی است.",
	},

	"CurrentPassword": {
		"min": "رمز عبور اصلی باید حداقل ۸ کاراکتر باشد.",
	},

	"NewPassword": {
		"min": "رمز عبور جدید باید حداقل ۸ کاراکتر باشد.",
	},
}
