package web

type DepositRequest struct {
	AccountNumber int `json:"account_number"`
	Nominal       int `json:"nominal"`
}

type DepositResponse struct {
	Balance int `json:"balance"`
}
