package ipapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type IpapiData struct {
	IP                 string  `json:"ip"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	Country            string  `json:"country"`
	CountryName        string  `json:"country_name"`
	CountryCode        string  `json:"country_code"`
	CountryCodeIso3    string  `json:"country_code_iso3"`
	CountryCapital     string  `json:"country_capital"`
	CountryTld         string  `json:"country_tld"`
	ContinentCode      string  `json:"continent_code"`
	InEu               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UtcOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	CountryArea        float64 `json:"country_area"`
	CountryPopulation  int     `json:"country_population"`
	Asn                string  `json:"asn"`
	Org                string  `json:"org"`
}

const (
	serviceUrl     = "https://ipapi.co/json/"
	allowedCountry = "Cyprus"
	HeaderKey      = "User-Agent"
	HeaderValue    = "ipapi.co/#go-v1.18"
)

func GetData() (*IpapiData, error) {
	data := &IpapiData{}
	ipapiClient := http.Client{}
	req, err := http.NewRequest("GET", serviceUrl, nil)
	req.Header.Set(HeaderKey, HeaderValue)
	resp, err := ipapiClient.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return data, nil
}

func IsAllowed() (bool, string) {
	data, err := GetData()
	if err != nil {
		return false, ""
	}

	return data.CountryName == allowedCountry, data.CountryName
}
