package Entities

type AccountsList struct {
	Count int       `json:"count"`
	Data  []Account `json:"data"`
}

type InstitutionsList struct {
	Count int           `json:"count"`
	Data  []Institution `json:"data"`
}

type Connection struct {
	Id          string           `json:"id"`
	Status      string           `json:"status"`
	LastUsed    string           `json:"lastUsed"`
	Institution AccountsList     `json:"institution"`
	Accounts    InstitutionsList `json:"accounts"`
}
