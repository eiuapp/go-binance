package futures

import (
	"context"
	"encoding/json"
)

// ListTopLongShortPositionRatioService trades
type ListTopLongShortPositionRatioService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *ListTopLongShortPositionRatioService) Symbol(symbol string) *ListTopLongShortPositionRatioService {
	s.symbol = symbol
	return s
}

// Period set period
func (s *ListTopLongShortPositionRatioService) Period(period string) *ListTopLongShortPositionRatioService {
	s.period = period
	return s
}

// Limit set limit
func (s *ListTopLongShortPositionRatioService) Limit(limit int) *ListTopLongShortPositionRatioService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *ListTopLongShortPositionRatioService) StartTime(startTime int64) *ListTopLongShortPositionRatioService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListTopLongShortPositionRatioService) EndTime(endTime int64) *ListTopLongShortPositionRatioService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *ListTopLongShortPositionRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*ListTopLongShortPositionRatio, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/futures/data/topLongShortPositionRatio",
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
	res = make([]*ListTopLongShortPositionRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return
}

// ListTopLongShortPositionRatio define open interest history info
type ListTopLongShortPositionRatio struct {
	Symbol         string `json:"symbol"`
	LongShortRatio string `json:"longShortRatio"`
	LongAccount    string `json:"longAccount"`
	ShortAccount   string `json:"shortAccount"`
	Timestamp      int64  `json:"timestamp"`
}

// NewListTopLongShortPositionRatioService 大户账户数多空比
func (c *Client) NewListTopLongShortPositionRatioService() *ListTopLongShortPositionRatioService {
	return &ListTopLongShortPositionRatioService{c: c}
}
