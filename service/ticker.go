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

			switch time.Now().Hour() {
			case 8:
				TX(60 * 60 * 24 * 2)
			case 12:
				TX(60 * 60 * 24 * 2)
			case 18:
				TX(60 * 60 * 24 * 2)
			case 22:
				TX(60 * 60 * 24 * 2)
			case 15:
				TX(60 * 60 * 24 * 2)
			default:
				TX(60 * 60 * 24 * 1)
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
