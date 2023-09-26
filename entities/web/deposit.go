package web

type DepositRequest struct {
	AccountNumber string `json:"account_number"`
	Nominal       int    `json:"nominal"`
}
