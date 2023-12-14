package mackerelremo

import (
	"context"
)

type RunnerConfig struct {
	MackerelAPIKey          string `env:"MACKEREL_API_KEY"`
	MackerelServiceName     string `env:"MACKEREL_SERVICE_NAME"`
	NatureAccessToken       string `env:"NATURE_ACCESS_TOKEN"`
	RemoDeviceID            string `env:"REMO_DEVICE_ID"`
	RemoDeviceNameForExport string `env:"REMO_DEVICE_NAME_FOR_EXPORT"`
}

type Runner struct {
	*RunnerConfig
}

func NewRunner(cfg *RunnerConfig) *Runner {
	return &Runner{cfg}
}

func (r *Runner) Run(ctx context.Context) error {
	fetcher := NewFetcher(&FetcherConfig{
		NatureAccessToken: r.NatureAccessToken,
		RemoDeviceID:      r.RemoDeviceID,
	})
	result, err := fetcher.Fetch(ctx)
	if err != nil {
		return err
	}
	exporter := NewMackerelExporter(&MackerelExporterConfig{
		MackerelApiKey:          r.MackerelAPIKey,
		MackerelServiceName:     r.MackerelServiceName,
		RemoDeviceNameForExport: r.RemoDeviceNameForExport,
	})
	return exporter.Export(ctx, result)
}
