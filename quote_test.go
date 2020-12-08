package yahoofinance

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRegularMarketPriceCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	yfinanceTest, _ := New(client)

	tests := []struct {
		name    string
		c       *RegularMarketPriceCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewQuoteService(yfinanceTest).RegularMarketPrice("0050.TW"),
			"https://query1.finance.yahoo.com/v8/finance/chart/0050.TW?includeAdjustedClose=false&interval=1d&range=1d", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("RegularMarketPriceCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegularMarketPriceCall.doRequest() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
func TestRegularMarketPriceCall_Do(t *testing.T) {
	str := `{
		"chart": {
		  "result": [
			{
			  "meta": {
				"currency": "USD",
				"symbol": "VTI",
				"exchangeName": "PCX",
				"instrumentType": "ETF",
				"firstTradeDate": 992611800,
				"regularMarketTime": 1607374800,
				"gmtoffset": -18000,
				"timezone": "EST",
				"exchangeTimezoneName": "America/New_York",
				"regularMarketPrice": 191.3,
				"chartPreviousClose": 55.665,
				"priceHint": 2,
				"currentTradingPeriod": {
				  "pre": {
					"timezone": "EST",
					"start": 1607331600,
					"end": 1607351400,
					"gmtoffset": -18000
				  },
				  "regular": {
					"timezone": "EST",
					"start": 1607351400,
					"end": 1607374800,
					"gmtoffset": -18000
				  },
				  "post": {
					"timezone": "EST",
					"start": 1607374800,
					"end": 1607389200,
					"gmtoffset": -18000
				  }
				},
				"dataGranularity": "1d",
				"range": "max",
				"validRanges": [
				  "1d",
				  "5d",
				  "1mo",
				  "3mo",
				  "6mo",
				  "1y",
				  "2y",
				  "5y",
				  "10y",
				  "ytd",
				  "max"
				]
			  },
			  "timestamp": [
				992611800,
				992871000,
				1607092200
			  ],
			  "events": {
				"dividends": {
				  "993475800": {
					"amount": 0.14,
					"date": 993475800
				  },
				  "1601040600": {
					"amount": 0.674,
					"date": 1601040600
				  }
				},
				"splits": {
				  "1213795800": {
					"date": 1213795800,
					"numerator": 2,
					"denominator": 1,
					"splitRatio": "2:1"
				  }
				}
			  },
			  "indicators": {
				"quote": [
				  {
					"volume": [
					  1067400,
					  282600,
					  4401400
					],
					"close": [
					  55.665000915527344,
					  55.310001373291016,
					  191.50999450683594
					],
					"open": [
					  55.42499923706055,
					  55.814998626708984,
					  190
					],
					"high": [
					  56.005001068115234,
					  55.915000915527344,
					  191.50999450683594
					],
					"low": [
					  55.17499923706055,
					  55.310001373291016,
					  189.99000549316406
					]
				  }
				],
				"adjclose": [
				  {
					"adjclose": [
					  38.816429138183594,
					  38.568904876708984,
					  191.50999450683594
					]
				  }
				]
			  }
			}
		  ],
		  "error": null
		}
	  }`
	client := clientTest(str, http.StatusOK)
	yfinanceTest, _ := New(client)

	tests := []struct {
		name    string
		c       *RegularMarketPriceCall
		want    *Infomation
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test", NewQuoteService(yfinanceTest).RegularMarketPrice("VTI"), &Infomation{
			ServerResponse: ServerResponse{
				HTTPStatusCode: 200,
				Header:         map[string][]string{},
			},
			Chart: Chart{
				Result: []Result{
					{
						Meta: Meta{
							Currency:             "USD",
							Symbol:               "VTI",
							ExchangeName:         "PCX",
							InstrumentType:       "ETF",
							FirstTradeDate:       992611800,
							RegularMarketTime:    1607374800,
							Gmtoffset:            -18000,
							Timezone:             "EST",
							ExchangeTimezoneName: "America/New_York",
							RegularMarketPrice:   191.3,
							ChartPreviousClose:   55.665,
							PriceHint:            2,
							CurrentTradingPeriod: CurrentTradingPeriod{
								Pre:     TimeInfo{Timezone: "EST", Start: 1607331600, End: 1607351400, Gmtoffset: -18000},
								Regular: TimeInfo{Timezone: "EST", Start: 1607351400, End: 1607374800, Gmtoffset: -18000},
								Post:    TimeInfo{Timezone: "EST", Start: 1607374800, End: 1607389200, Gmtoffset: -18000},
							},
							DataGranularity: "1d",
							Range:           "max",
							ValidRanges:     []string{"1d", "5d", "1mo", "3mo", "6mo", "1y", "2y", "5y", "10y", "ytd", "max"},
						},
						Timestamp: []int64{992611800, 992871000, 1607092200},
						Events: Events{
							Dividends: map[string]Dividend{
								"1601040600": {Amount: 0.674, Date: 1601040600},
								"993475800":  {Amount: 0.14, Date: 993475800},
							},
							Splits: map[string]Split{
								"1213795800": {Date: 1213795800, Numerator: 2, Denominator: 1, SplitRatio: "2:1"},
							},
						},
						Indicators: Indicators{
							Quote: []Quote{
								{
									Volume: []float64{1067400, 282600, 4401400},
									Close:  []float64{55.665000915527344, 55.310001373291016, 191.50999450683594},
									Open:   []float64{55.42499923706055, 55.814998626708984, 190},
									High:   []float64{56.005001068115234, 55.915000915527344, 191.50999450683594},
									Low:    []float64{55.17499923706055, 55.310001373291016, 189.99000549316406},
								},
							},
							Adjclose: []Adjclose{
								{[]float64{38.816429138183594, 38.568904876708984, 191.50999450683594}},
							},
						},
					}},
				Error: ErrorHistory{},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("PeriodCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PeriodCall.Do() = \n%+v, \nwant \n%+v", got, tt.want)
			}
		})
	}
}
