package opts

import "github.com/otetz/go-tachart/types"

func Bool(val bool) types.Bool {
	return &val
}

func Int(val int) types.Int {
	return &val
}

func Float(val float32) types.Float {
	return &val
}

func Str(val string) types.String {
	return types.String(val)
}
