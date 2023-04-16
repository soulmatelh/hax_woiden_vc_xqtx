package service

import "time"

var Tickers *time.Ticker

func Ticker() {
	Tickers = time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-Tickers.C:
			Run()
		}
	}
}
func TickerTX() {
	Tickers := time.NewTicker(1 * time.Hour)
	for {
		select {
		case <-Tickers.C:
			if time.Now().Hour() == 8 || time.Now().Hour() == 20 || time.Now().Hour() == 0 {
				TX(60 * 60 * 24 * 2)
			}
		}
	}
}

func Write() {
	WTickers := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-WTickers.C:
			_ = Date.WriteJson()
		}
	}
}
