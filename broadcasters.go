package mlr

import (
	"fmt"
	"strings"
)

const (
	TRN_MATCH_LINK_FORMAT = "https://www.therugbynetwork.com/videos/%s-%s-round-%d-mlr-%s-live"
)

type Broadcaster struct {
	Name            string
	MatchLinkFormat string
}

func (b *Broadcaster) MatchLink(match Match) string {
	awayTeamFormatted := strings.ToLower(strings.Join(strings.Split(match.AwayTeam.Name, " "), "-"))
	homeTeamFormatted := strings.ToLower(strings.Join(strings.Split(match.HomeTeam.Name, " "), "-"))

	if strings.TrimSpace(awayTeamFormatted) == "rugby-atl" {
		awayTeamFormatted = "atl-rugby"
	}

	if strings.TrimSpace(homeTeamFormatted) == "rubgy-atl" {
		homeTeamFormatted = "atl-rugby"
	}

	switch b.Name {
	case "The Rugby Network":
		return fmt.Sprintf(TRN_MATCH_LINK_FORMAT, awayTeamFormatted, homeTeamFormatted, match.RoundNumber, match.SeasonName)
	case "TRN":
		return fmt.Sprintf(TRN_MATCH_LINK_FORMAT, awayTeamFormatted, homeTeamFormatted, match.RoundNumber, match.SeasonName)
	}

	return ""
}
