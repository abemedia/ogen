// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
	"math/big"
	"net/http"
	"regexp"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/nonrecording"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
)

var regexMap = map[string]*regexp.Regexp{
	"^(\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?$": regexp.MustCompile("^(\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?$"),
	"^\\d-\\d$": regexp.MustCompile("^\\d-\\d$"),
	"foo.*":     regexp.MustCompile("foo.*"),
	"string_.*": regexp.MustCompile("string_.*"),
}
var ratMap = map[string]*big.Rat{
	"10/1": func() *big.Rat {
		r := new(big.Rat)
		if err := r.UnmarshalText([]byte("10/1")); err != nil {
			panic(fmt.Sprintf("rat %q: %v", "10/1", err))
		}
		return r
	}(),
	"5/1": func() *big.Rat {
		r := new(big.Rat)
		if err := r.UnmarshalText([]byte("5/1")); err != nil {
			panic(fmt.Sprintf("rat %q: %v", "5/1", err))
		}
		return r
	}(),
}

type config struct {
	TracerProvider trace.TracerProvider
	Tracer         trace.Tracer
	MeterProvider  metric.MeterProvider
	Meter          metric.Meter
	Client         ht.Client
}

func newConfig(opts ...Option) config {
	cfg := config{
		TracerProvider: otel.GetTracerProvider(),
		MeterProvider:  nonrecording.NewNoopMeterProvider(),
		Client:         http.DefaultClient,
	}
	for _, opt := range opts {
		opt.apply(&cfg)
	}
	cfg.Tracer = cfg.TracerProvider.Tracer(otelogen.Name,
		trace.WithInstrumentationVersion(otelogen.SemVersion()),
	)
	cfg.Meter = cfg.MeterProvider.Meter(otelogen.Name)
	return cfg
}

type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	o(c)
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
//
// If none is specified, the global provider is used.
func WithTracerProvider(provider trace.TracerProvider) Option {
	return optionFunc(func(cfg *config) {
		if provider != nil {
			cfg.TracerProvider = provider
		}
	})
}

// WithMeterProvider specifies a meter provider to use for creating a meter.
//
// If none is specified, the metric.NewNoopMeterProvider is used.
func WithMeterProvider(provider metric.MeterProvider) Option {
	return optionFunc(func(cfg *config) {
		if provider != nil {
			cfg.MeterProvider = provider
		}
	})
}

// WithClient specifies http client to use.
func WithClient(client ht.Client) Option {
	return optionFunc(func(cfg *config) {
		if client != nil {
			cfg.Client = client
		}
	})
}