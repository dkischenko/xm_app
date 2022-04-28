package company

type CompanyCreateResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Hash string `json:"hash"`
}
