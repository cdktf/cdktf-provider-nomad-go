// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package csivolume

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CsiVolumeTopologyRequestPreferredTopologyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CsiVolumeTopologyRequestPreferredTopologyList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CsiVolumeTopologyRequestPreferredTopologyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeTopologyRequestPreferredTopologyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeTopologyRequestPreferredTopologyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeTopologyRequestPreferredTopologyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeTopologyRequestPreferredTopologyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCsiVolumeTopologyRequestPreferredTopologyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

