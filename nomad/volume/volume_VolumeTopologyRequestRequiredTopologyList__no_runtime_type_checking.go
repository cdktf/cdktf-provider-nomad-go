//go:build no_runtime_type_checking

package volume

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VolumeTopologyRequestRequiredTopologyList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VolumeTopologyRequestRequiredTopologyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VolumeTopologyRequestRequiredTopologyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VolumeTopologyRequestRequiredTopologyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VolumeTopologyRequestRequiredTopologyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VolumeTopologyRequestRequiredTopologyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVolumeTopologyRequestRequiredTopologyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

