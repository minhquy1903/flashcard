package common

type Response struct {
	Data  interface{} `json:"data"`
	Total int32       `json:"total,omitempty"`
	Page  int32       `json:"page,omitempty"`
}
