package Models

type VerificationCode struct {
	Id               int    `json:"id"`
	VerificationCode string `json:"code"`
	UserEmail        string `json:"user_email"`
}
