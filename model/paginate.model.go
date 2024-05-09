package model

type ParamPaginate struct {
	Search    string
	Page      int
	Limit     int
	Offset    int
	OrderBy   string
	Direction string
}

type ResponsePaginate struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"total_pages"`
	TotalData  int64 `json:"total_data"`
}
