//go:build no_runtime_type_checking

package externalvolume

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalVolumeTopologiesList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ExternalVolumeTopologiesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeTopologiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeTopologiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeTopologiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewExternalVolumeTopologiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

