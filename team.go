package mlr

import "google.golang.org/api/calendar/v3"

var CalendarLink = map[string]string{
	"MajorLeagueRugby":                     "0kibhubi5a4bc659pjv08cuitc@group.calendar.google.com",
	"a5f160ed-07c0-4236-8d74-632342f72dd9": "883e5mpg4rcfpu865hcoodejug@group.calendar.google.com",                                       // New York Iron Workers
	"e7ce85d6-6f5e-43e6-a7cb-4d15d3c70e7c": "mf59f1djmf1uogl0s7df4044jc@group.calendar.google.com",                                       // Toronto Arrows
	"a5bd5b72-ff43-4861-9056-86df1b9c9c96": "lp01if4cvpggmmkqlfinuhsvcc@group.calendar.google.com",                                       // Dallas Jackals
	"6a4512c0-f54f-4979-942c-e823d99fa748": "0vfstp3rd7f0b7cgg9ppdplnpk@group.calendar.google.com",                                       // Houston SaberCats
	"2f4e6b3e-7d3c-4e15-bace-a414976216fb": "mv999j4826q3evpdcfe4hbv5do@group.calendar.google.com",                                       // Rugby ATL
	"034db172-942f-48b8-bc91-a0b3eb3a025f": "ph07rn6ucgagjj97l91agmopms@group.calendar.google.com",                                       // Seattle Seawolves
	"a8c0ea92-6d79-4774-95f2-5a558a2c70ca": "8fbns5t1g5kgsn56j13emr4ffg@group.calendar.google.com",                                       // NOLA Gold
	"35c03f68-8fc0-4e15-b7ab-f4d559520218": "mia75mmf4epc5c1tsdttcjsd44@group.calendar.google.com",                                       // New England Free Jacks
	"e1b2f662-5fc9-4990-b426-d14f0b80b953": "3fmpe4b3raaclaa7brvb8388a4@group.calendar.google.com",                                       // Utah Warriors
	"7ea9b87e-8a71-450d-a744-545048d7dc0d": "n902u1hbk10jp1ovf9mj7b97pc@group.calendar.google.com",                                       // San Diego Legion
	"17a788b5-2ac6-41d6-a320-f4d75cdd08b9": "f1fef7eb46f3c2e82b98290844de276e7e193f56d626573d99e496b46788568b@group.calendar.google.com", // Chicago Hounds
	"09f042b2-7a4f-4fa2-bfc2-21e7469c764a": "dtvl535ihoi0a53u69ec2o7lr0@group.calendar.google.com",                                       // Old Glory DC
}

type Team struct {
	GUID    string           `json:"teamGuid"`
	Name    string           `json:"name"`
	Color   string           `json:"color"`
	Image   string           `json:"imageGuid"`
	Active  bool             `json:"isActive"`
	Matches []calendar.Event `json:"matches"`
	// Not Implemented
	// MetricDTOs
}

func (t *Team) Calendar() string {
	return CalendarLink[t.GUID]
}
