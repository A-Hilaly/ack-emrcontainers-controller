package job_run

import (
	"encoding/json"

	svcsdk "github.com/aws/aws-sdk-go/service/emrcontainers"
)

func cfgToString(cfg *svcsdk.ConfigurationOverrides) (*string, error) {
	configBytes, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	configStr := string(configBytes)
	return &configStr, nil
}

func stringToConfigurationOverrides(cfg *string) (*svcsdk.ConfigurationOverrides, error) {
	var config svcsdk.ConfigurationOverrides
	err := json.Unmarshal([]byte(*cfg), &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
