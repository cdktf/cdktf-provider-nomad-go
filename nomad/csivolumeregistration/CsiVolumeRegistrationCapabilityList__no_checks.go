// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package csivolumeregistration

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CsiVolumeRegistrationCapabilityList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CsiVolumeRegistrationCapabilityList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CsiVolumeRegistrationCapabilityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeRegistrationCapabilityList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeRegistrationCapabilityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeRegistrationCapabilityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CsiVolumeRegistrationCapabilityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCsiVolumeRegistrationCapabilityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

