package web

type MutationsResponse struct {
	KodeTransaksi   string `json:"transaction_code"`
	TransactionTime string `json:"transaction_time"`
	Nominal         int    `json:"nominal"`
}
