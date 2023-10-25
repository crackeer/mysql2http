package define

// QueryRequest
type QueryRequest struct {
	Query    map[string]interface{} `json:"query"`
	Modify   map[string]interface{} `json:"modify"`
	OrderBy  string                 `json:"order_by"`
	Page     int64                  `json:"page"`
	PageSize int64                  `json:"page_size"`
}
