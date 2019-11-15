package handlers

//nolint
type RespData struct {
	Success     bool   `json:"success"`
	Data        []List `json:"data,omitempty"`
	Description string `json:"description,omitempty"`
}

type List struct {
	Network string
	IsWhite bool
}

type RespParams struct {
	Success     bool   `json:"success"`
	Data        Params `json:"data,omitempty"`
	Description string `json:"description,omitempty"`
}

type Params struct {
	K, M, N int32
}
