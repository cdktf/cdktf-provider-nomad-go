// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package csivolume

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CsiVolumeTopologiesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CsiVolumeTopologiesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CsiVolumeTopologiesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeTopologiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeTopologiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeTopologiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCsiVolumeTopologiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

