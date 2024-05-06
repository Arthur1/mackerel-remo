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
	values := make([]*mackerel.MetricValue, 0, 4)
	t := result.Timestamp.Round(time.Minute).Unix()
	if result.Temperature != nil {
		values = append(values, &mackerel.MetricValue{
			Name:  fmt.Sprintf("natureremo.temperature.%s", e.RemoDeviceNameForExport),
			Time:  t,
			Value: result.Temperature.Value,
		}, &mackerel.MetricValue{
			Name:  fmt.Sprintf("natureremo.temperature.event_delay_seconds.%s", e.RemoDeviceNameForExport),
			Time:  t,
			Value: result.Temperature.Delay.Seconds(),
		})
	}
	if result.Humidity != nil {
		values = append(values, &mackerel.MetricValue{
			Name:  fmt.Sprintf("natureremo.humidity.%s", e.RemoDeviceNameForExport),
			Time:  t,
			Value: result.Humidity.Value,
		}, &mackerel.MetricValue{
			Name:  fmt.Sprintf("natureremo.humidity.event_delay_seconds.%s", e.RemoDeviceNameForExport),
			Time:  t,
			Value: result.Humidity.Delay.Seconds(),
		})
	}
	return e.client.PostServiceMetricValues(e.MackerelServiceName, values)
}
