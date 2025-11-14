// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dynamichostvolumeregistration

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamicHostVolumeRegistrationConstraintList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DynamicHostVolumeRegistrationConstraintList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DynamicHostVolumeRegistrationConstraintList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DynamicHostVolumeRegistrationConstraintList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamicHostVolumeRegistrationConstraintList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DynamicHostVolumeRegistrationConstraintList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDynamicHostVolumeRegistrationConstraintListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

