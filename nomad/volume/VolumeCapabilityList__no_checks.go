//go:build no_runtime_type_checking

package volume

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VolumeCapabilityList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VolumeCapabilityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VolumeCapabilityList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VolumeCapabilityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VolumeCapabilityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VolumeCapabilityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVolumeCapabilityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

