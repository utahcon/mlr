package mlr

type Round struct {
	GUID   string `json:"roundGuid"`
	Number int    `json:"roundNumber"`
	Name   string `json:"roundName"`
}
