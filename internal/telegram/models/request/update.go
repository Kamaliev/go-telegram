package request

type Update struct {
	Offset  int64 `json:"offset,omitempty"`
	Limit   int   `json:"limit,omitempty"`
	Timeout int   `json:"timeout,omitempty"`
}
