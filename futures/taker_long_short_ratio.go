package futures

import (
	"context"
	"encoding/json"
)

// ListTakerLongShortRatioService trades
type ListTakerLongShortRatioService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *ListTakerLongShortRatioService) Symbol(symbol string) *ListTakerLongShortRatioService {
	s.symbol = symbol
	return s
}

// Period set period
func (s *ListTakerLongShortRatioService) Period(period string) *ListTakerLongShortRatioService {
	s.period = period
	return s
}

// Limit set limit
func (s *ListTakerLongShortRatioService) Limit(limit int) *ListTakerLongShortRatioService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *ListTakerLongShortRatioService) StartTime(startTime int64) *ListTakerLongShortRatioService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListTakerLongShortRatioService) EndTime(endTime int64) *ListTakerLongShortRatioService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *ListTakerLongShortRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*ListTakerLongShortRatio, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/futures/data/takerlongshortRatio",
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
	res = make([]*ListTakerLongShortRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return
}

// ListTakerLongShortRatio define open interest history info
type ListTakerLongShortRatio struct {
	BuySellRatio string `json:"buySellRatio"`
	BuyVol       string `json:"buyVol"`
	SellVol      string `json:"sellVol"`
	Timestamp    int64  `json:"timestamp"`
}

// NewListTakerLongShortRatioService 大户账户数多空比
func (c *Client) NewListTakerLongShortRatioService() *ListTakerLongShortRatioService {
	return &ListTakerLongShortRatioService{c: c}
}
