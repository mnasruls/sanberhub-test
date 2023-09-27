package web

type UpdateBalanceRequest struct {
	AccountNumber int `json:"account_number"`
	Nominal       int `json:"nominal"`
}

type UpdateBalanceResponse struct {
	Balance int `json:"balance"`
}
