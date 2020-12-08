package yahoofinance

import (
	"context"
	"net/http"
	"net/url"
)

// ServerResponse is embedded in each Do response and
// provides the HTTP status code and header sent by the server.
type ServerResponse struct {
	// HTTPStatusCode is the server's response status code. When using a
	// resource method's Do call, this will always be in the 2xx range.
	HTTPStatusCode int
	// Header contains the response header fields from the server.
	Header http.Header
}

// DefaultCall DefaultCall function
type DefaultCall struct {
	s         *Service
	urlParams url.Values
	ctx       context.Context
	header    http.Header
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *DefaultCall) Context(ctx context.Context) *DefaultCall {
	c.ctx = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *DefaultCall) Header() http.Header {
	if c.header == nil {
		c.header = make(http.Header)
	}
	return c.header
}

// ===============================================================================================================

// TimeInfo TimeInfo
type TimeInfo struct {
	Timezone  string `json:"timezone"`
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
	Gmtoffset int64  `json:"gmtoffset"`
}

// Dividend Dividend
type Dividend struct {
	Amount float64 `json:"amount"`
	Date   int64   `json:"date"`
}

// Split Split
type Split struct {
	Date        int64  `json:"date"`
	Numerator   int    `json:"numerator"`
	Denominator int    `json:"denominator"`
	SplitRatio  string `json:"splitRatio"`
}

// CurrentTradingPeriod CurrentTradingPeriod
type CurrentTradingPeriod struct {
	Pre     TimeInfo `json:"pre"`
	Regular TimeInfo `json:"regular"`
	Post    TimeInfo `json:"post"`
}

// Meta Meta
type Meta struct {
	Currency             string               `json:"currency"`
	Symbol               string               `json:"symbol"`
	ExchangeName         string               `json:"exchangeName"`
	InstrumentType       string               `json:"instrumentType"`
	FirstTradeDate       int64                `json:"firstTradeDate"`
	RegularMarketTime    int64                `json:"regularMarketTime"`
	Gmtoffset            int64                `json:"gmtoffset"`
	Timezone             string               `json:"timezone"`
	ExchangeTimezoneName string               `json:"exchangeTimezoneName"`
	RegularMarketPrice   float64              `json:"regularMarketPrice"`
	ChartPreviousClose   float64              `json:"chartPreviousClose"`
	PriceHint            int                  `json:"priceHint"`
	CurrentTradingPeriod CurrentTradingPeriod `json:"currentTradingPeriod"`
	DataGranularity      string               `json:"dataGranularity"`
	Range                string               `json:"range"`
	ValidRanges          []string             `json:"validRanges"`
}

// Events Events
type Events struct {
	Dividends map[string]Dividend `json:"dividends"`
	Splits    map[string]Split    `json:"splits"`
}

// Quote Quote
type Quote struct {
	Volume []float64 `json:"volume"`
	Close  []float64 `json:"close"`
	Open   []float64 `json:"open"`
	High   []float64 `json:"high"`
	Low    []float64 `json:"low"`
}

// Adjclose Adjclose
type Adjclose struct {
	Value []float64 `json:"adjclose"`
}

// Indicators Indicators
type Indicators struct {
	Quote    []Quote    `json:"quote"`
	Adjclose []Adjclose `json:"adjclose"`
}

// Result Result
type Result struct {
	Meta       Meta       `json:"meta"`
	Timestamp  []int64    `json:"timestamp"`
	Events     Events     `json:"events"`
	Indicators Indicators `json:"indicators"`
}

// ErrorHistory ErrorHistory
type ErrorHistory struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// Chart Chart
type Chart struct {
	Result []Result     `json:"result"`
	Error  ErrorHistory `json:"error"`
}

// Infomation Infomation
type Infomation struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`
	Chart          Chart `json:"chart"`
}
