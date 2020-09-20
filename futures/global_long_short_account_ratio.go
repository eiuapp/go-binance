package futures

import (
	"context"
	"encoding/json"
)

// ListGlobalLongShortAccountRatioService trades
type ListGlobalLongShortAccountRatioService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *ListGlobalLongShortAccountRatioService) Symbol(symbol string) *ListGlobalLongShortAccountRatioService {
	s.symbol = symbol
	return s
}

// Period set period
func (s *ListGlobalLongShortAccountRatioService) Period(period string) *ListGlobalLongShortAccountRatioService {
	s.period = period
	return s
}

// Limit set limit
func (s *ListGlobalLongShortAccountRatioService) Limit(limit int) *ListGlobalLongShortAccountRatioService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *ListGlobalLongShortAccountRatioService) StartTime(startTime int64) *ListGlobalLongShortAccountRatioService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListGlobalLongShortAccountRatioService) EndTime(endTime int64) *ListGlobalLongShortAccountRatioService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *ListGlobalLongShortAccountRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*ListGlobalLongShortAccountRatio, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/futures/data/globalLongShortAccountRatio",
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
	res = make([]*ListGlobalLongShortAccountRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return
}

// ListGlobalLongShortAccountRatio define open interest history info
type ListGlobalLongShortAccountRatio struct {
	Symbol         string `json:"symbol"`
	LongShortRatio string `json:"longShortRatio"`
	LongAccount    string `json:"longAccount"`
	ShortAccount   string `json:"shortAccount"`
	Timestamp      int64  `json:"timestamp"`
}

// NewListGlobalLongShortAccountRatioService 大户账户数多空比
func (c *Client) NewListGlobalLongShortAccountRatioService() *ListGlobalLongShortAccountRatioService {
	return &ListGlobalLongShortAccountRatioService{c: c}
}
