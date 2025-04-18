package opts

import "github.com/otetz/go-tachart/types"

type AngleAxis struct {
	PolarAxisBase
	Clockwise types.Bool `json:"clockwise,omitempty"`
}
