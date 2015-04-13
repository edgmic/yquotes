package yquotes

import (
	"github.com/doneland/test"
	"testing"
	"time"
)

// Get stock price data.
func TestGetStockPrice(t *testing.T) {
	symbol := "AAPL"
	stock, err := GetPrice(symbol)

	test.Ok(t, err)
	test.Equals(t, symbol, stock.Symbol)
}

// Get stock price history.
func TestStockPriceHistory(t *testing.T) {
	symbol := "AAPL"
	from, _ := time.Parse("2006-01-02", "2015-01-01")
	to, _ := time.Parse("2006-01-02", "2015-04-10")
	prices, err := GetDailyHistory(symbol, from, to)

	test.Ok(t, err)
	test.Equals(t, int(4), int(prices[0].Date.Month()))
	test.Equals(t, int(10), int(prices[0].Date.Day()))
	test.Equals(t, int(2015), int(prices[0].Date.Year()))
}

// Get stock price for one day.
func TestGetStockPriceForDay(t *testing.T) {
	symbol := "AAPL"
	date, _ := time.Parse("2006-01-02", "2015-04-09")
	price, err := GetPriceForDate(symbol, date)

	test.Ok(t, err)
	test.Equals(t, int(4), int(price.Date.Month()))
}
