package ddwrapper

import (
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

type ExtendedClientInterface interface {
	statsd.ClientInterface
	GaugeFixedRate(name string, value float64, tags []string) error
	CountFixedRate(name string, value int64, tags []string) error
	HistogramFixedRate(name string, value float64, tags []string) error
	DistributionFixedRate(name string, value float64, tags []string) error
	DecrFixedRate(name string, tags []string) error
	IncrFixedRate(name string, tags []string) error
	SetFixedRate(name string, value string, tags []string) error
	TimingFixedRate(name string, value time.Duration, tags []string) error
	TimeInMillisecondsFixedRate(name string, value float64, tags []string) error
}
