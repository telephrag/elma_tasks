package models

type Response struct {
	Percent int               `json:"percent"`
	Fails   []ReceivedDesired `json:"fails"`
}

type ReceivedDesired struct {
	OriginalResult interface{} `json:"OriginalResult"`
	ExternalResult interface{} `json:"ExternalResult"`
	DataSet        float64     `json:"DataSet"`
}
