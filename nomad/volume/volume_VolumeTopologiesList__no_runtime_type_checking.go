//go:build no_runtime_type_checking

package volume

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VolumeTopologiesList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VolumeTopologiesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VolumeTopologiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VolumeTopologiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VolumeTopologiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVolumeTopologiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

