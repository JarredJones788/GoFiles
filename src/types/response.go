package types

//GenericResponse - Simple response
type GenericResponse struct {
	Response bool `json:"response"`
}

//URLResponse - response with a Url
type URLResponse struct {
	Response bool   `json:"response"`
	URL      string `json:"url"`
}

//ReasonResponse - return response with a reason
type ReasonResponse struct {
	Response bool   `json:"response"`
	Reason   string `json:"reason"`
}
