package charts

import (
	"github.com/otetz/go-tachart/opts"
	"github.com/otetz/go-tachart/render"
	"github.com/otetz/go-tachart/types"
)

// Pie represents a pie chart.
type Pie struct {
	BaseConfiguration
}

// Type returns the chart type.
func (*Pie) Type() string { return types.ChartPie }

// NewPie creates a new pie chart.
func NewPie() *Pie {
	c := &Pie{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

// AddSeries adds new data sets.
func (c *Pie) AddSeries(name string, data []opts.PieData, options ...SeriesOpts) *Pie {
	series := SingleSeries{Name: name, Type: types.ChartPie, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Pie instance.
func (c *Pie) SetGlobalOptions(options ...GlobalOpts) *Pie {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// Validate validates the given configuration.
func (c *Pie) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
