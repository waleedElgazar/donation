package Models

type VerificationCode struct {
	Id int `json:"id"`
	VerificationCode string `json:"code"`
	USerEmail	string `json:"user_email"`
}
