package charts

import (
	"github.com/otetz/go-tachart/opts"
	"github.com/otetz/go-tachart/render"
	"github.com/otetz/go-tachart/types"
)

// Radar represents a radar chart.
type Radar struct {
	BaseConfiguration

	// SymbolKeepAspect is whether to keep aspect for symbols in the form of path://.
	SymbolKeepAspect types.Bool
}

// Type returns the chart type.
func (*Radar) Type() string { return types.ChartRadar }

// NewRadar creates a new radar chart.
func NewRadar() *Radar {
	c := &Radar{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.hasRadar = true
	return c
}

// AddSeries adds new data sets.
func (c *Radar) AddSeries(name string, data []opts.RadarData, options ...SeriesOpts) *Radar {
	series := SingleSeries{Name: name, Type: types.ChartRadar, Data: data, SymbolKeepAspect: c.SymbolKeepAspect}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	c.legends = append(c.legends, name)
	return c
}

// SetGlobalOptions sets options for the Radar instance.
func (c *Radar) SetGlobalOptions(options ...GlobalOpts) *Radar {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// Validate validates the given configuration.
func (c *Radar) Validate() {
	c.Legend.Data = c.legends
	c.Assets.Validate(c.AssetsHost)
}
