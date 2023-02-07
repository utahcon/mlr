package mlr

import (
	"google.golang.org/api/calendar/v3"
	"strings"
	"time"
)

type mlrTime struct {
	time.Time
}

const mlrTimeFormat = "2006-01-02T15:04:05"

func (m *mlrTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	m.Time, err = time.Parse(mlrTimeFormat, s)
	if err != nil {
		return err
	}
	return nil
}

type Match struct {
	GUID        string         `json:"matchGuid"`
	Season      string         `json:"seasonGuid,omitempty"`
	SeasonName  string         `json:"seasonName"`
	Series      string         `json:"series,omitempty"`
	SeriesName  string         `json:"seriesName"`
	Date        mlrTime        `json:"date"`
	UTCDate     mlrTime        `json:"utcDate"`
	RoundNumber int            `json:"roundNumber"`
	RoundName   string         `json:"roundName,omitempty"`
	PoolGUID    string         `json:"poolGuid,omitempty"`
	PoolName    string         `json:"poolName"`
	VenueGUID   string         `json:"VenueGuid"`
	Venue       Venue          `json:"venue,omitempty"`
	Details     string         `json:"matchDetails,omitempty"`
	HomeTeamId  string         `json:"homeTeamID,omitempty"`
	AwayTeamId  string         `json:"awayTeamID,omitempty"`
	AwayTeam    Team           `json:"awayTeam"`
	HomeTeam    Team           `json:"homeTeam"`
	TeamDTOs    []Team         `json:"teamDTOs"`
	Referees    []Referee      `json:"referees"`
	Event       calendar.Event `json:"event"`

	ScoreSummaryPublished bool   `json:"isScoreSummaryPublished"`
	MatchDataPublished    bool   `json:"isMatchDataPublished"`
	LiveCoded             bool   `json:"isLoveCoded"`
	MatchDataCoded        bool   `json:"isMatchDataCoded"`
	MatchDataValidated    bool   `json:"isMatchDataValidated"`
	Cancelled             bool   `json:"isCancelled"`
	Postponed             bool   `json:"isPostponed"`
	HomeForfeit           bool   `json:"isHomeForfeit"`
	AwayForfeit           bool   `json:"isAwayForfeit"`
	Complete              bool   `json:"isComplete"`
	TicketLink            string `json:"mlrTicketLink"`
	Broadcasters          string `json:"broadcasters"`

	// Not Implemented
	//MetricDTOs
}
