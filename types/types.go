package types

type Articles struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

type IDResponse struct {
	ID int64 `json:"id"`
}

type Response struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}
