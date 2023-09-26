package web

type RegisterRequest struct {
	Name string `json:"name"`
	NIK  string `json:"nik"`
	NoHP string `json:"no_hp"`
}

type RegistResponse struct {
	AccountNumber int `json:"account_number"`
}
