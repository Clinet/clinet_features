package features

import (
	"github.com/Clinet/clinet_cmds"
	"github.com/Clinet/clinet_convos"
	"github.com/Clinet/clinet_services"
)

type FeatureMap struct {
	Features map[string]Feature `json:"features"`
}
var FM *FeatureMap

type Feature struct {
	Toggle       bool                `json:"toggle,omitempty"`
	Name         string              `json:"name"`
	Desc         string              `json:"-"`
	Init         func() error        `json:"-"`
	Cmds         []*cmds.Cmd         `json:"-"`
	ServiceChat  services.Service    `json:"-"`
	ServiceConvo convos.ConvoService `json:"-"`
}

func SetFeatures(features map[string]Feature) {
	if FM == nil {
		FM = &FeatureMap{}
	}
	FM.Features = features
}
func IsEnabled(feature string) bool {
	f, exists := FM.Features[feature]
	if exists {
		return f.Toggle
	}
	return false
}
func UpdateFeature(feature Feature) {
	if FM != nil && FM.Features != nil {
		if _, exists := FM.Features[feature.Name]; exists {
			feature.Toggle = FM.Features[feature.Name].Toggle
			FM.Features[feature.Name] = feature
		}
	}
}
func GetEnabledFeatures() []string {
	enabledFeatures := make([]string, 0)
	for fName, f := range FM.Features {
		if f.Toggle {
			enabledFeatures = append(enabledFeatures, fName)
		}
	}
	return enabledFeatures
}