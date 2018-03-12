package Utilities

type API struct {
	host string
}

func NewAPI(host string) *API {
	return &API{
		host: host,
	}
}
