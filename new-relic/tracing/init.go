package newrelictracing

import (
	"context"

	"github.com/newrelic/go-agent/v3/integrations/nrlogrus"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ServiceName                string
	LicenseKey                 string
	IsDistributedTracerEnabled bool
	IsLoggingEnabled           bool
}

var App *newrelic.Application

// Init - Initialize App
func Init(ctx context.Context, config Config) error {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.ServiceName),
		newrelic.ConfigLicense(config.LicenseKey),
		newrelic.ConfigDistributedTracerEnabled(config.IsDistributedTracerEnabled),
		func(nrConfig *newrelic.Config) {
			if config.IsLoggingEnabled == true {
				logrus.SetLevel(logrus.DebugLevel)
				nrConfig.Logger = nrlogrus.StandardLogger()
			}
		},
	)

	if err != nil {
		return err
	}

	App = app
	return nil
}
