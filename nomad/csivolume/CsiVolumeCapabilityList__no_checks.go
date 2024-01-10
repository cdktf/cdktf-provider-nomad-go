// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package csivolume

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CsiVolumeCapabilityList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CsiVolumeCapabilityList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CsiVolumeCapabilityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeCapabilityList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeCapabilityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeCapabilityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeCapabilityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCsiVolumeCapabilityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

