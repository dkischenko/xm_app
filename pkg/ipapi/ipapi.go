package ipapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

type IpapiData struct {
	CountryName string `json:"country_name"`
}

const (
	serviceUrl     = "https://ipapi.co/"
	responseType   = "/json"
	allowedCountry = "Cyprus"
	HeaderKey      = "User-Agent"
	HeaderValue    = "ipapi.co/#go-v1.18"
)

func GetData(ip string) (*IpapiData, error) {
	data := &IpapiData{}
	ipapiClient := http.Client{}
	url := serviceUrl + ip + responseType
	req, err := http.NewRequest("GET", url, nil)
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

func IsAllowed(addr string) (bool, string) {
	ip, _, err := net.SplitHostPort(addr)
	if err != nil {
		return false, ""
	}
	data, err := GetData(ip)
	if err != nil {
		return false, ""
	}

	return data.CountryName == allowedCountry, data.CountryName
}
