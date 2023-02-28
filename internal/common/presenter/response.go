package dto

type Response struct {
	Data  interface{} `json:"data"`
	Total int32       `json:"total,omitempty"`
	Page  int32       `json:"page,omitempty"`
}

type GetListResponse struct {
	Items      interface{} `json:"items"`
	TotalCount int64       `json:"total_count"`
	PerPage    int         `json:"per_page"`
	Page       int         `json:"page"`
}
