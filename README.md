# yahoofinance - Yahoo Finance API in Go
[![GoDoc](https://godoc.org/github.com/z-Wind/yahoofinance?status.png)](http://godoc.org/github.com/z-Wind/yahoofinance)

## Table of Contents

* [Installation](#installation)
* [Examples](#examples)
* [Reference](#reference)

## Installation

    $ go get github.com/z-Wind/yahoofinance

## Examples

### Client
```go
client := GetClient()
yfinance, err := New(client)
```

### History
```go
call := yfinance.History.Period("0050.TW", "1mo", "1d")
history, err := call.Do()
```


## Reference
- [https://github.com/ranaroussi/yfinance](https://github.com/ranaroussi/yfinance)
