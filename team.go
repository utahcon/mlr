package mlr

import "google.golang.org/api/calendar/v3"

//0kibhubi5a4bc659pjv08cuitc@group.calendar.google.com: Major League Rugby
//lp01if4cvpggmmkqlfinuhsvcc@group.calendar.google.com: MLR: Dallas Jackals
//0vfstp3rd7f0b7cgg9ppdplnpk@group.calendar.google.com: MLR: Houston SaberCats
//mia75mmf4epc5c1tsdttcjsd44@group.calendar.google.com: MLR: New England Free Jacks
//8fbns5t1g5kgsn56j13emr4ffg@group.calendar.google.com: MLR: NOLA Gold
//dtvl535ihoi0a53u69ec2o7lr0@group.calendar.google.com: MLR: Old Glory DC
//mv999j4826q3evpdcfe4hbv5do@group.calendar.google.com: MLR: Rugby ATL
//883e5mpg4rcfpu865hcoodejug@group.calendar.google.com: MLR: Rugby New York
//n902u1hbk10jp1ovf9mj7b97pc@group.calendar.google.com: MLR: San Diego Legion
//ph07rn6ucgagjj97l91agmopms@group.calendar.google.com: MLR: Seattle Seawolves
//mf59f1djmf1uogl0s7df4044jc@group.calendar.google.com: MLR: Toronto Arrows
//3fmpe4b3raaclaa7brvb8388a4@group.calendar.google.com: MLR: Utah Warriors

type Team struct {
	GUID   string `json:"teamGuid"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	Image  string `json:"imageGuid"`
	Active bool   `json:"isActive"`
	//MetricDTOs
	CalendarId string           `json:"calendar-id"`
	Matches    []calendar.Event `json:"matches"`
}
