package mackerelremo

import (
	"context"
	"fmt"
	"time"

	"github.com/mackerelio/mackerel-client-go"
)

type MackerelExporter struct {
	*MackerelExporterConfig
	client *mackerel.Client
}

type MackerelExporterConfig struct {
	MackerelApiKey          string
	MackerelServiceName     string
	RemoDeviceNameForExport string
}

func NewMackerelExporter(cfg *MackerelExporterConfig) *MackerelExporter {
	cli := mackerel.NewClient(cfg.MackerelApiKey)
	return &MackerelExporter{cfg, cli}
}

func (e *MackerelExporter) Export(ctx context.Context, result *FetchResult) error {
	values := make([]*mackerel.MetricValue, 0, 2)
	if result.Temperature != nil {
		values = append(values, &mackerel.MetricValue{
			Name:  fmt.Sprintf("natureremo.temperature.%s", e.RemoDeviceNameForExport),
			Time:  result.Temperature.Timestamp.Round(time.Minute).Unix(),
			Value: result.Temperature.Value,
		})
	}
	if result.Humidity != nil {
		values = append(values, &mackerel.MetricValue{
			Name:  fmt.Sprintf("natureremo.humidity.%s", e.RemoDeviceNameForExport),
			Time:  result.Humidity.Timestamp.Round(time.Minute).Unix(),
			Value: result.Humidity.Value,
		})
	}
	return e.client.PostServiceMetricValues(e.MackerelServiceName, values)
}
