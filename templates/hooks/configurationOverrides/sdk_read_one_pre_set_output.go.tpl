// DescribeJobRun should output ConfigurationOverrides and show all available configuration
if resp.JobRun.ConfigurationOverrides != nil {
	configBytes, err := json.Marshal(resp.JobRun.ConfigurationOverrides)
	if err != nil {
		return err
	}
	ko.Spec.Configuration = configBytes
}
