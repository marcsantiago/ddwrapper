package ddwrapper

import (
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

var _ ExtendedClientInterface = (*Client)(nil)

type Client struct {
	client *statsd.Client
}

// Gauge measures the value of a metric at a particular time.
func (c *Client) Gauge(name string, value float64, tags []string, rate float64) error {
	return c.client.Gauge(name, value, Tags(tags).Pair(), rate)
}

// GaugeFixedRate measures the value of a metric at a particular time. Rate is fixed to 1.
func (c *Client) GaugeFixedRate(name string, value float64, tags []string) error {
	return c.client.Gauge(name, value, Tags(tags).Pair(), 1)
}

// Count tracks how many times something happened per second.
func (c *Client) Count(name string, value int64, tags []string, rate float64) error {
	return c.client.Count(name, value, Tags(tags).Pair(), rate)
}

// CountFixedRate tracks how many times something happened per second. Rate is fixed to 1.
func (c *Client) CountFixedRate(name string, value int64, tags []string) error {
	return c.client.Count(name, value, Tags(tags).Pair(), 1)
}

// Histogram tracks the statistical distribution of a set of values on each host.
func (c *Client) Histogram(name string, value float64, tags []string, rate float64) error {
	return c.client.Histogram(name, value, Tags(tags).Pair(), rate)
}

// HistogramFixedRate tracks the statistical distribution of a set of values on each host. Rate is fixed to 1.
func (c *Client) HistogramFixedRate(name string, value float64, tags []string) error {
	return c.client.Histogram(name, value, Tags(tags).Pair(), 1)
}

// Distribution tracks the statistical distribution of a set of values across your infrastructure.
func (c *Client) Distribution(name string, value float64, tags []string, rate float64) error {
	return c.client.Distribution(name, value, Tags(tags).Pair(), rate)
}

// DistributionFixedRate tracks the statistical distribution of a set of values across your infrastructure. Rate is fixed to 1.
func (c *Client) DistributionFixedRate(name string, value float64, tags []string) error {
	return c.client.Distribution(name, value, Tags(tags).Pair(), 1)
}

// Decr is just Count of -1
func (c *Client) Decr(name string, tags []string, rate float64) error {
	return c.client.Decr(name, Tags(tags).Pair(), rate)
}

// DecrFixedRate is just Count of -1. Rate is fixed to 1.
func (c *Client) DecrFixedRate(name string, tags []string) error {
	return c.client.Decr(name, Tags(tags).Pair(), 1)
}

// Incr is just Count of 1
func (c *Client) Incr(name string, tags []string, rate float64) error {
	return c.client.Incr(name, Tags(tags).Pair(), rate)
}

// IncrFixedRate is just Count of 1. Rate is fixed to 1.
func (c *Client) IncrFixedRate(name string, tags []string) error {
	return c.client.Incr(name, Tags(tags).Pair(), 1)
}

// Set counts the number of unique elements in a group.
func (c *Client) Set(name string, value string, tags []string, rate float64) error {
	return c.client.Set(name, value, Tags(tags).Pair(), rate)
}

// SetFixedRate counts the number of unique elements in a group. Rate is fixed to 1.
func (c *Client) SetFixedRate(name string, value string, tags []string) error {
	return c.client.Set(name, value, Tags(tags).Pair(), 1)
}

// Timing sends timing information, it is an alias for TimeInMilliseconds
func (c *Client) Timing(name string, value time.Duration, tags []string, rate float64) error {
	return c.client.Timing(name, value, Tags(tags).Pair(), rate)
}

// TimingFixedRate sends timing information, it is an alias for TimeInMilliseconds. Rate is fixed to 1.
func (c *Client) TimingFixedRate(name string, value time.Duration, tags []string) error {
	return c.client.Timing(name, value, Tags(tags).Pair(), 1)
}

// TimeInMilliseconds sends timing information in milliseconds.
func (c *Client) TimeInMilliseconds(name string, value float64, tags []string, rate float64) error {
	return c.client.TimeInMilliseconds(name, value, Tags(tags).Pair(), rate)
}

// TimeInMillisecondsFixedRate sends timing information in milliseconds. Rate is fixed to 1.
func (c *Client) TimeInMillisecondsFixedRate(name string, value float64, tags []string) error {
	return c.client.TimeInMilliseconds(name, value, Tags(tags).Pair(), 1)
}

// Event sends the provided Event.
func (c *Client) Event(e *statsd.Event) error {
	return c.client.Event(e)
}

// SimpleEvent sends an event with the provided title and text.
func (c *Client) SimpleEvent(title, text string) error {
	return c.client.SimpleEvent(title, text)
}

// ServiceCheck sends the provided ServiceCheck.
func (c *Client) ServiceCheck(sc *statsd.ServiceCheck) error {
	return c.client.ServiceCheck(sc)
}

// SimpleServiceCheck sends an serviceCheck with the provided name and status.
func (c *Client) SimpleServiceCheck(name string, status statsd.ServiceCheckStatus) error {
	return c.client.SimpleServiceCheck(name, status)
}

// Close the client connection.
// Acquire closer lock to ensure only one thread can close the stop channel
// Notify all other threads that they should stop
// Wait for the threads to stop
// Finally flush any remaining metrics that may have come in at the last moment
func (c *Client) Close() error {
	return c.Close()
}

// Flush forces a flush of all the queued dogstatsd payloads
// This method is blocking and will not return until everything is sent
// through the network
func (c *Client) Flush() error {
	return c.client.Flush()
}

// SetWriteTimeout allows the user to set a custom write timeout.
func (c *Client) SetWriteTimeout(dur time.Duration) error {
	return c.client.SetWriteTimeout(dur)
}
