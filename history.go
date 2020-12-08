package yahoofinance

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// NewHistoryService get history
func NewHistoryService(s *Service) *HistoryService {
	rs := &HistoryService{s: s}
	return rs
}

// HistoryService get history
type HistoryService struct {
	s *Service
}

// Period get data in period
// https://query1.finance.yahoo.com/v8/finance/chart/0050.TW?range=7d&interval=1d&includeAdjustedClose=true&events="div,splits"
// https://query1.finance.yahoo.com/v8/finance/chart/VTI?period1=-2208988800&period2=1607299200&interval=1d&includeAdjustedClose=true&events="div,splits"
/*
   :Parameters:
       period : str
           Valid periods: 1d,5d,1mo,3mo,6mo,1y,2y,5y,10y,ytd,max
           Either Use period parameter or use start and end
       interval : str
           Valid intervals: 1m,2m,5m,15m,30m,60m,90m,1h,1d,5d,1wk,1mo,3mo
           Intraday data cannot extend last 60 days
       start: str
           Download start date string (YYYY-MM-DD) or _datetime.
           Default is 1900-01-01
       end: str
           Download end date string (YYYY-MM-DD) or _datetime.
           Default is now
*/
func (r *HistoryService) Period(symbol, period, interval string) *PeriodCall {
	c := &PeriodCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		symbol: symbol,
	}

	c.urlParams.Set("range", strings.ToLower(period))
	c.urlParams.Set("interval", strings.ToLower(interval))
	c.urlParams.Set("includeAdjustedClose", "true")
	c.urlParams.Set("events", "div,splits")

	return c
}

// PeriodCall call function
type PeriodCall struct {
	DefaultCall

	symbol string
}

// IncludeAdjustedClose Adjust Close Default is true
func (c *PeriodCall) IncludeAdjustedClose(s string) *PeriodCall {
	c.urlParams.Set("includeAdjustedClose", strings.ToLower(s))
	return c
}

func (c *PeriodCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	reqHeaders.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	// 無需設定 http.Transport 已自帶，並自動解碼，若加上會產生亂碼
	// reqHeaders.Set("Accept-Encoding", "gzip, deflate, br")
	reqHeaders.Set("Accept-Language", "zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	var body io.Reader = nil
	urls := ResolveRelative(c.s.host, "/v8/finance/chart", c.symbol)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *PeriodCall) Do() (*Infomation, error) {
	res, err := c.doRequest()
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, errors.Wrapf(err, "doRequest")
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	ret := &Infomation{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}

	return ret, nil
}

// Between get data in period
// https://query1.finance.yahoo.com/v8/finance/chart/0050.TW?range=7d&interval=1d&includeAdjustedClose=true&events="div,splits"
// https://query1.finance.yahoo.com/v8/finance/chart/VTI?period1=-2208988800&period2=1607299200&interval=1d&includeAdjustedClose=true&events="div,splits"
/*
   :Parameters:
       period : str
           Valid periods: 1d,5d,1mo,3mo,6mo,1y,2y,5y,10y,ytd,max
           Either Use period parameter or use start and end
       interval : str
           Valid intervals: 1m,2m,5m,15m,30m,60m,90m,1h,1d,5d,1wk,1mo,3mo
           Intraday data cannot extend last 60 days
       start: str
           Download start date string (YYYY-MM-DD) or _datetime.
           Default is 1900-01-01
       end: str
           Download end date string (YYYY-MM-DD) or _datetime.
           Default is now
*/
func (r *HistoryService) Between(symbol string, start, end time.Time) *BetweenCall {
	c := &BetweenCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		symbol: symbol,
	}

	c.urlParams.Set("period1", strconv.FormatInt(start.Unix(), 10))
	c.urlParams.Set("period2", strconv.FormatInt(end.Unix(), 10))
	c.urlParams.Set("interval", "1d")
	c.urlParams.Set("includeAdjustedClose", "true")
	c.urlParams.Set("events", "div,splits")

	return c
}

// BetweenCall call function
type BetweenCall struct {
	DefaultCall

	symbol string
}

// Interval Default is 1d
// Valid intervals: 1m,2m,5m,15m,30m,60m,90m,1h,1d,5d,1wk,1mo,3mo
// Intraday data cannot extend last 60 days
func (c *BetweenCall) Interval(s string) *BetweenCall {
	c.urlParams.Set("interval", strings.ToLower(s))
	return c
}

// IncludeAdjustedClose Adjust Close Default is true
func (c *BetweenCall) IncludeAdjustedClose(s string) *BetweenCall {
	c.urlParams.Set("includeAdjustedClose", strings.ToLower(s))
	return c
}

func (c *BetweenCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	reqHeaders.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	// 無需設定 http.Transport 已自帶，並自動解碼，若加上會產生亂碼
	// reqHeaders.Set("Accept-Encoding", "gzip, deflate, br")
	reqHeaders.Set("Accept-Language", "zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7")

	var body io.Reader = nil
	urls := ResolveRelative(c.s.host, "/v8/finance/chart", c.symbol)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *BetweenCall) Do() (*Infomation, error) {
	res, err := c.doRequest()
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, errors.Wrapf(err, "doRequest")
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	ret := &Infomation{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}

	return ret, nil
}
