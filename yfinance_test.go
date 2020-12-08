package yahoofinance

import (
	"fmt"
	"testing"
)

func TestNewServer(t *testing.T) {
	client := GetClient()
	yfinance, err := New(client)
	if err != nil {
		t.Fatal(err)
	}

	call := yfinance.History.Period("0050.TW", "1mo", "1d")
	history, err := call.Do()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", history)

	call = yfinance.History.Period("0050.W", "1mo", "1d")
	history, err = call.Do()
	if err == nil {
		t.Fatal("Should be Fail")
	}
}

func ExampleHistoryService_Period() {
	client := GetClient()
	yfinance, err := New(client)
	if err != nil {
		panic(err)
	}

	call := yfinance.History.Period("0050.TW", "1mo", "1d")
	history, err := call.Do()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", history)
}
