package charts

import (
	"github.com/otetz/go-tachart/datasets"
	"github.com/otetz/go-tachart/opts"
	"github.com/otetz/go-tachart/render"
	"github.com/otetz/go-tachart/types"
)

// Map represents a map chart.
type Map struct {
	BaseConfiguration

	mapType string
}

// Type returns the chart type.
func (*Map) Type() string { return types.ChartMap }

// NewMap creates a new map chart.
func NewMap() *Map {
	c := &Map{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

// RegisterMapType registers the given mapType.
func (c *Map) RegisterMapType(mapType string) {
	c.mapType = mapType
	mapFile, preset := datasets.PresetMapFileNames[mapType]
	if preset {
		c.JSAssets.Add("maps/" + mapFile + ".js")
	}
}

// AddSeries adds new data sets.
func (c *Map) AddSeries(name string, data []opts.MapData, options ...SeriesOpts) *Map {
	series := SingleSeries{Name: name, Type: types.ChartMap, MapType: c.mapType, Data: data}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Map instance.
func (c *Map) SetGlobalOptions(options ...GlobalOpts) *Map {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// Validate validates the given configuration.
func (c *Map) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
