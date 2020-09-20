package futures

import (
	"context"
	"encoding/json"
)

// OpenInterestHistService trades
type OpenInterestHistService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *OpenInterestHistService) Symbol(symbol string) *OpenInterestHistService {
	s.symbol = symbol
	return s
}

// Period set period
func (s *OpenInterestHistService) Period(period string) *OpenInterestHistService {
	s.period = period
	return s
}

// Limit set limit
func (s *OpenInterestHistService) Limit(limit int) *OpenInterestHistService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *OpenInterestHistService) StartTime(startTime int64) *OpenInterestHistService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *OpenInterestHistService) EndTime(endTime int64) *OpenInterestHistService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *OpenInterestHistService) Do(ctx context.Context, opts ...RequestOption) (res []*OpenInterestHist, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/futures/data/openInterestHist",
		secType:  secTypeAPIKey,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("period", s.period)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res = make([]*OpenInterestHist, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return
}

// OpenInterestHist define open interest history info
type OpenInterestHist struct {
	Symbol               string `json:"symbol"`
	SumOpenInterest      string `json:"sumOpenInterest"`
	SumOpenInterestValue string `json:"sumOpenInterestValue"`
	Timestamp            int64  `json:"timestamp"`
}

// NewOpenInterestHistService init list open interest history service
func (c *Client) NewOpenInterestHistService() *OpenInterestHistService {
	return &OpenInterestHistService{c: c}
}
