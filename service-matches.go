package mlr

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type MatchesService struct {
	s *Service
}

type MatchListCall struct {
	s            *Service
	urlParams_   URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

type MatchList struct {
	Matches  []Match
	Response *http.Response
}

func NewMatchesService(s *Service) *MatchesService {
	rs := &MatchesService{s: s}
	return rs
}

func (r *MatchesService) List() *MatchListCall {
	c := &MatchListCall{
		s:          r.s,
		urlParams_: make(URLParams),
		header_:    http.Header{},
	}
	c.header_.Add("Host", "stratusmedia.mobii.com")
	c.header_.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:108.0) Gecko/20100101 Firefox/108.0")
	c.header_.Add("Accept", "application/json")
	c.header_.Add("Accept-Language", "en-US,en;q=0.5")
	c.header_.Add("Accept-Encoding", "gzip, deflate, br")
	c.header_.Add("Access-Control-Request-Method", "GET")
	c.header_.Add("Access-Control-Request-Headers", "content-type,editonbehalfofuserguid,stratus-media-headers,tenantid")
	c.header_.Add("Content-Type", "application/json")
	c.header_.Add("Referer", "https://www.majorleague.rugby/")
	c.header_.Add("Origin", "https://www.majorleague.rugby")
	c.header_.Add("DNT", "1")
	c.header_.Add("Connection", "keep-alive")
	c.header_.Add("Sec-Fetch-Dest", "empty")
	c.header_.Add("Sec-Fetch-Mode", "cors")
	c.header_.Add("Sec-Fetch-Site", "cross-site")
	c.header_.Add("Sec-GPC", "1")
	c.header_.Add("Pragma", "no-cache")
	c.header_.Add("Cache-Control", "no-cache")
	c.header_.Add("TE", "trailers")
	c.header_.Add("tenantId", "4CFEC22F-C6D7-4BDC-A568-B1071DBD0B6E")
	c.header_.Add("EditOnBehalfOfUserGuid", "1EFF3C3F-1069-44FE-9B87-17F22B57597F")
	c.header_.Add("stratus-media-headers", "true")
	return c
}

func (c *MatchListCall) TeamOne(team string) *MatchListCall {
	c.urlParams_.Set("searchTeamOneId", team)
	return c
}

func (c *MatchListCall) TeamTwo(team string) *MatchListCall {
	c.urlParams_.Set("searchTeamTwoId", team)
	return c
}

func (c *MatchListCall) Series(series string) *MatchListCall {
	c.urlParams_.Set("seriesId", series)
	return c
}

func (c *MatchListCall) Season(season string) *MatchListCall {
	c.urlParams_.Set("seasonId", season)
	return c
}

func (c *MatchListCall) Count(size int) *MatchListCall {
	c.urlParams_.Set("pageSize", strconv.Itoa(size))
	return c
}

func (c *MatchListCall) Page(page int) *MatchListCall {
	c.urlParams_.Set("pageNumber", strconv.Itoa(page))
	return c
}

func (c *MatchListCall) ExcludePlayers(excludePlayers bool) *MatchListCall {
	c.urlParams_.Set("excludePlayers", fmt.Sprintf("%t", excludePlayers))
	return c
}

func (c *MatchListCall) Do() (*MatchList, error) {
	resp, err := c.doRequest()
	if err != nil {
		return nil, err
	}

	var matchList []Match

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.Unmarshal(b, &matchList)
	if err != nil {
		return nil, err
	}

	matchListResult := &MatchList{
		Matches:  matchList,
		Response: resp,
	}

	return matchListResult, nil
}

func (c *MatchListCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-mlr-api-client", Version)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	urls := SearchMatches
	urls += "?" + c.urlParams_.Encode()

	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return SendRequest(c.s.client, req)
}
