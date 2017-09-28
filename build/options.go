package build

import "encoding/json"

type Options struct {
	Root          string
	BinName       string
	ExtractAssets bool
	HasDB         bool
	LDFlags       string
	Tags          []string
	Static        bool
	Debug         bool
	Compress      bool
}

func (o Options) String() string {
	b, _ := json.Marshal(o)
	return string(b)
}
