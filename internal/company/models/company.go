package models

type Company struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Code      int    `json:"code"`
	CountryId int    `json:"countryId"`
	Website   string `json:"website"`
	Phone     string `json:"phone"`
}
