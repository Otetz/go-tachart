package charts

import (
	"github.com/otetz/go-tachart/opts"
	"github.com/otetz/go-tachart/render"
	"github.com/otetz/go-tachart/types"
)

// Tree represents a Tree chart.
type Tree struct {
	BaseConfiguration
}

// Type returns the chart type.
func (*Tree) Type() string { return types.ChartTree }

// NewTree creates a new Tree chart instance.
func NewTree() *Tree {
	c := &Tree{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

// AddSeries adds new data sets.
func (c *Tree) AddSeries(name string, data []opts.TreeData, options ...SeriesOpts) *Tree {
	series := SingleSeries{Name: name, Type: types.ChartTree, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Tree instance.
func (c *Tree) SetGlobalOptions(options ...GlobalOpts) *Tree {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// Validate validates the given configuration.
func (c *Tree) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
