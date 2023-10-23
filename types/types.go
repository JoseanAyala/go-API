package types

type Articles struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
	Description string `json:"description"`
}

type IDResponse struct {
	ID int64 `json:"id"`
}
