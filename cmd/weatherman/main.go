package main

import (
	"fmt"
	"os"

	"github.com/Maliud/golang-Weather/internal/httpclient"
	"github.com/Maliud/golang-Weather/internal/logger"
	"github.com/Maliud/golang-Weather/pkg/client"
)


func main() {
	logger := logger.NewLogger(true)
	logger.Debug("hava Durumu Başlatılıyor...")

	httpClient := httpclient.NewHttpClient()

	cl, err := client.NewClient(logger, httpClient)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	res, err := cl.GetWeather(5.999802, 56.308920)
	if err != nil {
		logger.Debug("GetWeather() getırırken hata olustu")
		logger.Error(err)
		os.Exit(1)
	}
	logger.Debug("db'den yanıt geldi!")
	fmt.Printf(res)
}