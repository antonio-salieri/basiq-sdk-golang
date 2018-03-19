package services

type TransactionsList struct {
	Count int           `json:"count"`
	Data  []Transaction `json:"data"`
}

type Transaction struct {
	Id              string                 `json:"id"`
	Status          string                 `json:"status"`
	Description     string                 `json:"description"`
	PostDate        string                 `json:"postDate"`
	TransactionDate string                 `json:"transactionDate"`
	Amount          string                 `json:"amount"`
	Balance         string                 `json:"balance"`
	BankCategory    string                 `json:"bankCategory"`
	Account         string                 `json:"account"`
	Institution     string                 `json:"institution"`
	Connection      string                 `json:"connection"`
	Class           map[string]interface{} `json:"class"`
}
