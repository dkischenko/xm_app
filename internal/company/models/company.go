package models

type Company struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Code      int    `json:"code"`
	CountryId int    `json:"countryId"`
	Website   string `json:"website"`
	Phone     string `json:"phone"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

type CompanyCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Code    int    `json:"code" validate:"required,numeric"`
	Country string `json:"country" validate:"required,alpha"`
	Website string `json:"website" validate:"required,url"`
	Phone   string `json:"phone" validate:"required,e164"`
}

type CompanyUpdateRequest struct {
	Name      string `json:"name"`
	Code      int    `json:"code" validate:"numeric"`
	CountryId int    `json:"country_id" validate:"alpha"`
	Website   string `json:"website" validate:"url"`
	Phone     string `json:"phone" validate:"e164"`
}
