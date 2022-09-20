// Unmarshall ConfigurationOverrides structure
if desired.ko.Spec.Configuration != nil {
  var config svcsdk.ConfigurationOverrides
  err := json.Unmarshal(desired.ko.Spec.Configuration, &config)
  if err != nil {
    return err
  }
  input.ConfigurationOverrides = &config
}
