package charts

import (
	"github.com/otetz/go-tachart/opts"
	"github.com/otetz/go-tachart/render"
	"github.com/otetz/go-tachart/types"
)

// Parallel represents a parallel axis.
type Parallel struct {
	BaseConfiguration
}

// Type returns the chart type.
func (*Parallel) Type() string { return types.ChartParallel }

// NewParallel creates a new parallel instance.
func NewParallel() *Parallel {
	c := &Parallel{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasParallel = true
	return c
}

// AddSeries adds new data sets.
func (c *Parallel) AddSeries(name string, data []opts.ParallelData, options ...SeriesOpts) *Parallel {
	series := SingleSeries{Name: name, Type: types.ChartParallel, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Parallel instance.
func (c *Parallel) SetGlobalOptions(options ...GlobalOpts) *Parallel {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// Validate validates the given configuration.
func (c *Parallel) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
