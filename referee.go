package mlr

type Referee struct {
	RefereeGuid string `json:"refereeGuid"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DisplayName string `json:"displayName"`
	TypeID      string `json:"typeID"`
	TypeName    string `json:"typeName"`
	CountryCode string `json:"countryCode"`
	Image       string `json:"image"`
	//Metrics     []DTOs `json:"metricDTOs"`
}
