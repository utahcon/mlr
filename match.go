package mlr

import (
	"google.golang.org/api/calendar/v3"
	"time"
)

type Match struct {
	GUID         string         `json:"matchGuid"`
	Season       Season         `json:"seasonGuid,omitempty"`
	Round        Round          `json:"roundGuid,omitempty"`
	Series       Series         `json:"series,omitempty"`
	Pool         string         `json:"poolGuid,omitempty"`
	Date         time.Time      `json:"date"`
	UTCDate      time.Time      `json:"utcDate"`
	Venue        Venue          `json:"venue,omitempty"`
	Details      string         `json:"matchDetails,omitempty"`
	HomeTeam     Team           `json:"homeTeamID,omitempty"`
	AwayTeam     Team           `json:"awayTeamID,omitempty"`
	Cancelled    bool           `json:"isCancelled"`
	Postponed    bool           `json:"isPostponed"`
	HomeForfeit  bool           `json:"isHomeForfeit"`
	AwayForfeit  bool           `json:"isAwayForfeit"`
	Complete     bool           `json:"isComplete"`
	TicketLink   string         `json:"mlrTicketLink"`
	Broadcasters string         `json:"broadcasters"`
	Event        calendar.Event `json:"event"`
	Guest        Team           `json:"guest"`
	//Referees []Referee         `json:"referees"`
}
