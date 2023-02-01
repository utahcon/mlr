package mlr

type Venue struct {
	GUID       string `json:"venueGuid,omitempty"`
	Name       string `json:"name"`
	GoalWidth  int    `json:"goalWidth"`
	FieldWidth int    `json:"fieldWidth"`
	Active     bool   `json:"isActive"`
	Image      string `json:"image"`
}
