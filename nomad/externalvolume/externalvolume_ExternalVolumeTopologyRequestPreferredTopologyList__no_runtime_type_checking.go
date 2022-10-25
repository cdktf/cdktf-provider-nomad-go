//go:build no_runtime_type_checking

package externalvolume

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalVolumeTopologyRequestPreferredTopologyList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ExternalVolumeTopologyRequestPreferredTopologyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeTopologyRequestPreferredTopologyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeTopologyRequestPreferredTopologyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeTopologyRequestPreferredTopologyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ExternalVolumeTopologyRequestPreferredTopologyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewExternalVolumeTopologyRequestPreferredTopologyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

