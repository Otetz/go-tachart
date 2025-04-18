package components

import (
	"github.com/otetz/go-tachart/opts"
	"github.com/otetz/go-tachart/render"
)

// Charter represents a chart value which provides its type, assets and can be validated.
type Charter interface {
	Type() string
	GetAssets() opts.Assets
	// FillDefaultValues fill default values and would be overridden if any struct fields has been manually set
	FillDefaultValues()
	// Validate a validator as well as a post processor before render
	Validate()
}

// Page represents a page chart.
type Page struct {
	render.Renderer
	opts.Initialization
	opts.Assets

	PageBackgroundColor string
	Layout              Layout
	Charts              []interface{}
	ChartArea           ChartLayout
}

// NewPage creates a new page.
func NewPage(assetsHost string) *Page {
	page := &Page{}
	page.Assets.InitAssets()

	page.Renderer = render.NewPageRender(page, page.Validate)
	page.AssetsHost = assetsHost
	page.ChartArea = PageCenterLayout
	page.PageBackgroundColor = "#FFFFFF"

	return page
}

// SetLayout sets the layout of the Page.
func (page *Page) SetLayout(layout Layout) *Page {
	page.Layout = layout
	return page
}

func (page *Page) SetBackgroundColor(color string) *Page {
	page.PageBackgroundColor = color
	return page
}

// AddCharts adds new charts to the page and merge assets.
func (page *Page) AddCharts(charts ...Charter) *Page {
	for i := 0; i < len(charts); i++ {
		assets := charts[i].GetAssets()
		for _, v := range assets.JSAssets.Values {
			page.JSAssets.Add(v)
		}

		for _, v := range assets.CSSAssets.Values {
			page.CSSAssets.Add(v)
		}

		for _, v := range assets.CustomizedJSAssets.Values {
			page.CustomizedJSAssets.Add(v)
		}

		for _, v := range assets.CustomizedCSSAssets.Values {
			page.CustomizedCSSAssets.Add(v)
		}

		for _, v := range assets.CustomizedHeaders.Values {
			page.CustomizedHeaders.Add(v)
		}

		charts[i].Validate()
		page.Charts = append(page.Charts, charts[i])
	}
	return page
}

// Validate validates the given configuration.
func (page *Page) Validate() {
	page.Initialization.Validate()
	page.Assets.Validate(page.AssetsHost)
}
