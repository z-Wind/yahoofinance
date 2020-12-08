package yahoofinance

import (
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// NewQuoteService get history
func NewQuoteService(s *Service) *QuoteService {
	rs := &QuoteService{s: s}
	return rs
}

// QuoteService get history
type QuoteService struct {
	s *Service
}

// RegularMarketPrice get last price
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
func (r *QuoteService) RegularMarketPrice(symbol string) *RegularMarketPriceCall {
	c := &RegularMarketPriceCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		symbol: symbol,
	}

	c.urlParams.Set("range", "1d")
	c.urlParams.Set("interval", "1d")
	c.urlParams.Set("includeAdjustedClose", "false")

	return c
}

// RegularMarketPriceCall call function
type RegularMarketPriceCall struct {
	DefaultCall

	symbol string
}

func (c *RegularMarketPriceCall) doRequest() (*http.Response, error) {
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
func (c *RegularMarketPriceCall) Do() (*Infomation, error) {
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
