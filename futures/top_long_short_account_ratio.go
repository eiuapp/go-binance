package futures

import (
	"context"
	"encoding/json"
)

// ListTopLongShortAccountRatioService trades
type ListTopLongShortAccountRatioService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *ListTopLongShortAccountRatioService) Symbol(symbol string) *ListTopLongShortAccountRatioService {
	s.symbol = symbol
	return s
}

// Period set period
func (s *ListTopLongShortAccountRatioService) Period(period string) *ListTopLongShortAccountRatioService {
	s.period = period
	return s
}

// Limit set limit
func (s *ListTopLongShortAccountRatioService) Limit(limit int) *ListTopLongShortAccountRatioService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *ListTopLongShortAccountRatioService) StartTime(startTime int64) *ListTopLongShortAccountRatioService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListTopLongShortAccountRatioService) EndTime(endTime int64) *ListTopLongShortAccountRatioService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *ListTopLongShortAccountRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*ListTopLongShortAccountRatio, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/futures/data/topLongShortAccountRatio",
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
	res = make([]*ListTopLongShortAccountRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return
}

// ListTopLongShortAccountRatio define open interest history info
type ListTopLongShortAccountRatio struct {
	Symbol         string `json:"symbol"`
	LongShortRatio string `json:"longShortRatio"`
	LongAccount    string `json:"longAccount"`
	ShortAccount   string `json:"shortAccount"`
	Timestamp      int64  `json:"timestamp"`
}

// NewListTopLongShortAccountRatioService 大户账户数多空比
func (c *Client) NewListTopLongShortAccountRatioService() *ListTopLongShortAccountRatioService {
	return &ListTopLongShortAccountRatioService{c: c}
}
