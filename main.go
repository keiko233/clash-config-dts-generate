package main

import (
	"github.com/gzuidhof/tygo/tygo"
)

func main() {
	config := &tygo.Config{
		Packages: []*tygo.PackageConfig{
			{
				Path: "github.com/metacubex/mihomo/config",
				TypeMappings: map[string]string{
					"interface{}": "any /* hyperparameters */",
				},
				FallbackType: "any",
				OutputPath:   "./dts",
				Flavor:       "json",
			},
		},
	}

	gen := tygo.New(config)
	if err := gen.Generate(); err != nil {
		print(err.Error())
	}
}
