package charts

import (
	"github.com/otetz/go-tachart/opts"
	"github.com/otetz/go-tachart/render"
	"github.com/otetz/go-tachart/types"
)

// TreeMap represents a TreeMap chart.
type TreeMap struct {
	BaseConfiguration
}

// Type returns the chart type.
func (*TreeMap) Type() string { return types.ChartTreeMap }

// NewTreeMap creates a new TreeMap chart instance.
func NewTreeMap() *TreeMap {
	c := &TreeMap{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

// AddSeries adds new data sets.
func (c *TreeMap) AddSeries(name string, data []opts.TreeMapNode, options ...SeriesOpts) *TreeMap {
	series := SingleSeries{Name: name, Type: types.ChartTreeMap, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the TreeMap instance.
func (c *TreeMap) SetGlobalOptions(options ...GlobalOpts) *TreeMap {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// Validate validates the given configuration.
func (c *TreeMap) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
