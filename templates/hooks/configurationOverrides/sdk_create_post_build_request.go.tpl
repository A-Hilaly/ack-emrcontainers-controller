    // Unmarshall ConfigurationOverrides structure
    if desired.ko.Spec.Configuration != nil {
        input.ConfigurationOverrides, err = stringToConfigurationOverrides(desired.ko.Spec.Configuration)
        if err != nil {
          return nil, err
        }
    }