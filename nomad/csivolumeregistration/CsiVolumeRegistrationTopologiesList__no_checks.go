// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package csivolumeregistration

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CsiVolumeRegistrationTopologiesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CsiVolumeRegistrationTopologiesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CsiVolumeRegistrationTopologiesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeRegistrationTopologiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeRegistrationTopologiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeRegistrationTopologiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCsiVolumeRegistrationTopologiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

