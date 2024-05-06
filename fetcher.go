package mackerelremo

import (
	"context"
	"fmt"
	"time"

	"github.com/tenntenn/natureremo"
)

type FetcherConfig struct {
	NatureAccessToken string
	RemoDeviceID      string
}

type Fetcher struct {
	*FetcherConfig
	client *natureremo.Client
}

func NewFetcher(cfg *FetcherConfig) *Fetcher {
	cli := natureremo.NewClient(cfg.NatureAccessToken)
	return &Fetcher{cfg, cli}
}

type FetchResult struct {
	Timestamp   time.Time
	Temperature *FetchResultRow
	Humidity    *FetchResultRow
}

type FetchResultRow struct {
	Value float64
	Delay time.Duration
}

func (f *Fetcher) Fetch(ctx context.Context) (*FetchResult, error) {
	devices, err := f.client.DeviceService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var device *natureremo.Device
	for _, d := range devices {
		if d.ID == f.RemoDeviceID {
			device = d
			break
		}
	}
	if device == nil {
		return nil, fmt.Errorf("device not found")
	}
	now := time.Now()
	result := &FetchResult{
		Timestamp: now,
	}
	tempEvent, ok := device.NewestEvents[natureremo.SensorTypeTemperature]
	if ok {
		result.Temperature = &FetchResultRow{
			Value: tempEvent.Value,
			// tempEvent.CreatedAt will not be updated when Value does not change.
			Delay: now.Sub(tempEvent.CreatedAt),
		}
	}
	humidEvent, ok := device.NewestEvents[natureremo.SensorTypeHumidity]
	if ok {
		result.Humidity = &FetchResultRow{
			Value: humidEvent.Value,
			// humidEvent.CreatedAt will not be updated when Value does not change.
			Delay: now.Sub(humidEvent.CreatedAt),
		}
	}
	return result, nil
}
