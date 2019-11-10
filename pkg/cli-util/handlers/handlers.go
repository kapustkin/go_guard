package handlers

//nolint
type respData struct {
	Success     bool        `json:"success"`
	Data        interface{} `json:"data,omitempty"`
	Description string      `json:"description,omitempty"`
}
