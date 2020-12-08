package yahoofinance

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// const strings
const (
	// UserAgent is the header string used to identify this package.
	userAgent = `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:62.0) Gecko/20100101 Firefox/62.0`

	HOST = "https://query1.finance.yahoo.com/"
)

// Service Yahoo Finance api
type Service struct {
	client *http.Client

	host string // API endpoint base URL

	History *HistoryService
	Quote *QuoteService
}

// GetClient get client
func GetClient() *http.Client {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout:   0,
				KeepAlive: 0,
			}).Dial,
			TLSHandshakeTimeout: 10 * time.Second,
		},
		Jar: cookieJar,
	}

	return client
}

// New Yahoo Finance API server
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, host: HOST}
	s.History = NewHistoryService(s)
	s.Quote = NewQuoteService(s)

	return s, nil
}

func (s *Service) userAgent() string {
	return userAgent
}
