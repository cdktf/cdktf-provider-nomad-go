//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NomadProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (n *jsiiProxy_NomadProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateNomadProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_NomadProvider) validateSetHeadersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NomadProvider) validateSetIgnoreEnvVarsParameters(val *map[string]interface{}) error {
	return nil
}

func validateNewNomadProviderParameters(scope constructs.Construct, id *string, config *NomadProviderConfig) error {
	return nil
}

