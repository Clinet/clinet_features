package features

import (
	"github.com/Clinet/clinet_cmds"
	"github.com/Clinet/clinet_convos"
	"github.com/Clinet/clinet_services"
)

type FeatureMap struct {
	Features []Feature `json:"features"`
}
var FM *FeatureMap

type Feature struct {
	Toggle       bool                `json:"toggle"`
	Name         string              `json:"name"`
	Help         string              `json:"-"`
	Init         func() error        `json:"-"`
	Cmds         []*cmds.Cmd         `json:"-"`
	ServiceChat  services.Service    `json:"-"`
	ServiceConvo convos.ConvoService `json:"-"`
}

func SetFeatures(features []Feature) {
	if FM == nil {
		FM = &FeatureMap{}
	}
	FM.Features = features
}
func IsEnabled(feature string) bool {
	for _, f := range FM.Features {
		if feature == f.Name {
			return f.Toggle
		}
	}
	return false
}
func UpdateFeature(feature Feature) {
	if FM != nil {
		if FM.Features != nil && len(FM.Features) > 0 {
			for i := 0; i < len(FM.Features); i++ {
				if FM.Features[i].Name == feature.Name {
					feature.Toggle = FM.Features[i].Toggle
					FM.Features[i] = feature
				}
			}
		}
	}
}
