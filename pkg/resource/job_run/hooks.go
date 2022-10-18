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
	a *resource,
	b *resource,
) {
	aConfig, err := stringToConfigurationOverrides(a.ko.Spec.ConfigurationOverrides)
	if err != nil {
		panic(err)
	}
	bConfig, err := stringToConfigurationOverrides(b.ko.Spec.ConfigurationOverrides)
	if err != nil {
		panic(err)
	}

	if ackcompare.HasNilDifference(a.ko.Spec.ConfigurationOverrides, b.ko.Spec.ConfigurationOverrides) {
		delta.Add("Spec.ConfigurationOverrides", a.ko.Spec.ConfigurationOverrides, b.ko.Spec.ConfigurationOverrides)
	} else if a.ko.Spec.ConfigurationOverrides != nil && b.ko.Spec.ConfigurationOverrides != nil {
		delta.Add("Spec.ConfigurationOverrides", a.ko.Spec.ConfigurationOverrides, b.ko.Spec.ConfigurationOverrides)
	}

	if ackcompare.HasNilDifference(aConfig.ApplicationConfiguration, bConfig.ApplicationConfiguration) {
		delta.Add("Spec.ConfigurationOverrides", aConfig.ApplicationConfiguration, bConfig.ApplicationConfiguration)
	} else if aConfig.ApplicationConfiguration != nil && bConfig.ApplicationConfiguration != nil {
		delta.Add("Spec.ConfigurationOverrides", aConfig.ApplicationConfiguration, bConfig.ApplicationConfiguration)
	}
	else if aConfig.ApplicationConfiguration == nil && *bConfig.ApplicationConfiguration == "null" {
		delta.Add("Spec.ConfigurationOverrides", aConfig.ApplicationConfiguration, bConfig.ApplicationConfiguration)
	}

	if ackcompare.HasNilDifference(aConfig.MonitoringConfiguration, bConfig.MonitoringConfiguration) {
		delta.Add("Spec.ConfigurationOverrides", aConfig.MonitoringConfiguration, bConfig.MonitoringConfiguration)
	} else if aConfig.MonitoringConfiguration != nil && bConfig.MonitoringConfiguration !=nil {
		if ackcompare.HasNilDifference(aConfig.MonitoringConfiguration.PersistentAppUI, bConfig.MonitoringConfiguration.PersistentAppUI) {
			delta.Add("Spec.ConfigurationOverrides", aConfig.MonitoringConfiguration.PersistentAppUI, bConfig.MonitoringConfiguration.PersistentAppUI)
		} else if aConfig.MonitoringConfiguration.PersistentAppUI != nil && bConfig.MonitoringConfiguration.PersistentAppUI != nil {
			delta.Add("Spec.ConfigurationOverrides", aConfig.MonitoringConfiguration.PersistentAppUI, bConfig.MonitoringConfiguration.PersistentAppUI)
		} else if aConfig.MonitoringConfiguration.PersistentAppUI == nil && *bConfig.MonitoringConfiguration.PersistentAppUI == "ENABLED" {
			delta.Add("Spec.ConfigurationOverrides", aConfig.MonitoringConfiguration.PersistentAppUI, bConfig.MonitoringConfiguration.PersistentAppUI)
		}
	} else if ackcompare.HasNilDifference(aConfig.MonitoringConfiguration.CloudWatchMonitoringConfiguration, bConfig.MonitoringConfiguration.CloudWatchMonitoringConfiguration) {
			delta.Add("Spec.ConfigurationOverrides", aConfig.MonitoringConfiguration.CloudWatchMonitoringConfiguration, bConfig.MonitoringConfiguration.CloudWatchMonitoringConfiguration)
		} else if aConfig.MonitoringConfiguration.CloudWatchMonitoringConfiguration != nil && bConfig.MonitoringConfiguration.CloudWatchMonitoringConfiguration != nil {
			delta.Add("Spec.ConfigurationOverrides", aConfig.MonitoringConfiguration.CloudWatchMonitoringConfiguration, bConfig.MonitoringConfiguration.CloudWatchMonitoringConfiguration)
	} else if ackcompare.HasNilDifference(aConfig.MonitoringConfiguration.S3MonitoringConfiguration, bConfig.MonitoringConfiguration.S3MonitoringConfiguration) {
			delta.Add("Spec.ConfigurationOverrides", aConfig.MonitoringConfiguration.S3MonitoringConfiguration, bConfig.MonitoringConfiguration.S3MonitoringConfiguration)
		} else if aConfig.MonitoringConfiguration.S3MonitoringConfiguration != nil && bConfig.MonitoringConfiguration.S3MonitoringConfiguration != nil {
			delta.Add("Spec.ConfigurationOverrides", aConfig.MonitoringConfiguration.S3MonitoringConfiguration, bConfig.MonitoringConfiguration.S3MonitoringConfiguration)
	}
}
