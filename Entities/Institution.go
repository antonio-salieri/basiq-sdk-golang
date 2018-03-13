package Entities

type Institution struct {
	Id              string                 `json:"id"`
	Name            string                 `json:"name"`
	ShortName       string                 `json:"shortName"`
	Country         string                 `json:"country"`
	ServiceName     string                 `json:"serviceName"`
	ServiceType     string                 `json:"serviceType"`
	LoginIdCaption  string                 `json:"loginIdCaption"`
	PasswordCaption string                 `json:"PasswordCaption"`
	Colors          map[string]interface{} `json:"colors"`
	Logo            map[string]interface{} `json:"logo"`
}
