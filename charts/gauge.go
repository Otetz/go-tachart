package charts

import (
	"github.com/otetz/go-tachart/opts"
	"github.com/otetz/go-tachart/render"
	"github.com/otetz/go-tachart/types"
)

// Gauge represents a gauge chart.
type Gauge struct {
	BaseConfiguration
}

// Type returns the chart type.
func (*Gauge) Type() string { return types.ChartGauge }

// NewGauge creates a new gauge chart.
func NewGauge() *Gauge {
	c := &Gauge{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

// AddSeries adds new data sets.
func (c *Gauge) AddSeries(name string, data []opts.GaugeData, options ...SeriesOpts) *Gauge {
	series := SingleSeries{Name: name, Type: types.ChartGauge, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Gauge instance.
func (c *Gauge) SetGlobalOptions(options ...GlobalOpts) *Gauge {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// Validate validates the given configuration.
func (c *Gauge) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
