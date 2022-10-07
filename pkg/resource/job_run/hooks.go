package job_run

import (
	//"encoding/json"
	// "fmt"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	svcsdk "github.com/aws/aws-sdk-go/service/emrcontainers"
	"github.com/ghodss/yaml"
)

func cfgToString(cfg *svcsdk.ConfigurationOverrides) (*string, error) {
	configBytes, err := yaml.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	configStr := string(configBytes)
	return &configStr, nil
}

func stringToConfigurationOverrides(cfg *string) (*svcsdk.ConfigurationOverrides, error) {
	var config svcsdk.ConfigurationOverrides
	err := yaml.Unmarshal([]byte(*cfg), &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func customPreCompare(
	delta *ackcompare.Delta,
	// cfg *svcsdk.ConfigurationOverrides,
	a *resource,
	b *resource,
) {
	if ackcompare.HasNilDifference(a.ko.Spec.ConfigurationOverrides, b.ko.Spec.ConfigurationOverrides) {
		delta.Add("Spec.ConfigurationOverrides", a.ko.Spec.ConfigurationOverrides, b.ko.Spec.ConfigurationOverrides)
	} else if a.ko.Spec.ConfigurationOverrides != nil && b.ko.Spec.ConfigurationOverrides != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ConfigurationOverrides.ApplicationConfiguration, b.ko.Spec.ConfigurationOverrides.ApplicationConfiguration) {
			delta.Add("Spec.ConfigurationOverrides.ApplicationConfiguration", a.ko.Spec.ConfigurationOverrides.ApplicationConfiguration, b.ko.Spec.ConfigurationOverrides.ApplicationConfiguration)
		} else if a.ko.Spec.ConfigurationOverrides.ApplicationConfiguration != nil && b.ko.Spec.ConfigurationOverrides.ApplicationConfiguration != nil {
			if *a.ko.Spec.ConfigurationOverrides.ApplicationConfiguration != *b.ko.Spec.ConfigurationOverrides.ApplicationConfiguration {
				delta.Add("Spec.ConfigurationOverrides.ApplicationConfiguration", a.ko.Spec.ConfigurationOverrides.ApplicationConfiguration, b.ko.Spec.ConfigurationOverrides.ApplicationConfiguration)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ConfigurationOverrides.MonitoringConfiguration, b.ko.Spec.ConfigurationOverrides.MonitoringConfiguration) {
			delta.Add("Spec.ConfigurationOverrides.MonitoringConfiguration", a.ko.Spec.ConfigurationOverrides.MonitoringConfiguration, b.ko.Spec.ConfigurationOverrides.MonitoringConfiguration)
		} else if a.ko.Spec.ConfigurationOverrides.MonitoringConfiguration != nil && b.ko.Spec.ConfigurationOverrides.MonitoringConfiguration != nil {
			if *a.ko.Spec.ConfigurationOverrides.MonitoringConfiguration != *b.ko.Spec.ConfigurationOverrides.MonitoringConfiguration {
				delta.Add("Spec.ConfigurationOverrides.MonitoringConfiguration", a.ko.Spec.ConfigurationOverrides.MonitoringConfiguration, b.ko.Spec.ConfigurationOverrides.MonitoringConfiguration)
			}
		}
	}
}
