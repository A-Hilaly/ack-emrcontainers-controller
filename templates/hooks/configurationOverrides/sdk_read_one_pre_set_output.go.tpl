	if resp.JobRun.ConfigurationOverrides != nil {
		ko.Spec.Configuration, err = cfgToString(resp.JobRun.ConfigurationOverrides)
		if err != nil {
			return nil, err
		}
	}