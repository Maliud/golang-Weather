package client

import (
	"errors"
	"fmt"
)

type HttpClient interface {
	GetWeatherData(longitude float64, latitude float64) (string, error)
}

type Logger interface {
	Debug(msg string)
	Error(err error)
}

type Client struct {
	logger     Logger
	httpClient HttpClient
}

func NewClient(logger Logger, httpClient HttpClient) (*Client, error) {
	if logger == nil {
		return nil, errors.New("NewClient() çağırıldı ancak Logger nil oldu")
	}

	if httpClient == nil {
		return nil, errors.New("NewClient() çağırıldı ancak HttpClient nil oldu")
	}

	return &Client{
		logger: logger,
		httpClient: httpClient,
	}, nil
}

func (c *Client) GetWeather(longitude float64, latitude float64) (string, error) {
	c.logger.Debug(fmt.Sprintf(
		"GetWeather() %f %f ile çağrıldı. Oracle sor...", longitude, latitude,
	))

	return c.httpClient.GetWeatherData(longitude, latitude)
}
